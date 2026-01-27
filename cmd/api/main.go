package main

import (
	"log"

	"github.com/RafaelCarvalhoxd/financial-mangement/internal/auth"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/config"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/database"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/http/server"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env:", err)
	}

	cfg := config.NewConfig()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	userRepo := user.NewRepository(db)
	authService := auth.NewService(userRepo)
	authHandler := auth.NewHandler(authService, cfg)

	deps := &server.Dependencies{
		AuthHandler: authHandler,
	}
	router := server.Config(deps)
	port := ":" + cfg.Port

	log.Printf("Iniciando servidor na porta %s...", port)

	if err := router.Run(port); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
