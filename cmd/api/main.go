package main

import (
	"log"

	"github.com/RafaelCarvalhoxd/financial-management/internal/auth"
	"github.com/RafaelCarvalhoxd/financial-management/internal/category"
	"github.com/RafaelCarvalhoxd/financial-management/internal/http"
	"github.com/RafaelCarvalhoxd/financial-management/internal/infra/config"
	"github.com/RafaelCarvalhoxd/financial-management/internal/infra/database"
	"github.com/RafaelCarvalhoxd/financial-management/internal/transaction"
	"github.com/RafaelCarvalhoxd/financial-management/internal/user"
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

	transactionRepo := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := transaction.NewHandler(transactionService)

	deps := &http.Dependencies{
		AuthHandler:        authHandler,
		CategoryHandler:    categoryHandler,
		TransactionHandler: transactionHandler,
	}
	router := http.Config(deps)
	port := ":" + cfg.Port

	log.Printf("Iniciando servidor na porta %s...", port)

	if err := router.Run(port); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
