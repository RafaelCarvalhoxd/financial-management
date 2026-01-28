package main

import (
	"log"

	"github.com/RafaelCarvalhoxd/financial-management/internal/apps/auth"
	"github.com/RafaelCarvalhoxd/financial-management/internal/apps/category"
	"github.com/RafaelCarvalhoxd/financial-management/internal/apps/user"
	"github.com/RafaelCarvalhoxd/financial-management/internal/config"
	"github.com/RafaelCarvalhoxd/financial-management/internal/database"
	"github.com/RafaelCarvalhoxd/financial-management/internal/http/server"
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
	authService := auth.NewService(userRepo, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	categoryRepo := category.NewRepository(db)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := category.NewHandler(categoryService)

	deps := &server.Dependencies{
		AuthHandler:     authHandler,
		CategoryHandler: categoryHandler,
	}
	router := server.Config(deps)
	port := ":" + cfg.Port

	log.Printf("Iniciando servidor na porta %s...", port)

	if err := router.Run(port); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
