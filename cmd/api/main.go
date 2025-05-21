package main

import (
	"log"

	"insidechurch/internal/handlers"
	"insidechurch/internal/middleware"
	"insidechurch/internal/repositories"
	"insidechurch/internal/routes"
	"insidechurch/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configuração do banco de dados
	dsn := "host=localhost user=postgres password=postgres dbname=insidechurch port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Inicialização dos componentes
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	userHandler := handlers.NewUserHandler(userRepo, authService)
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Configuração do router
	router := gin.Default()

	// Configuração das rotas
	routes.SetupRoutes(router, userHandler, authMiddleware)

	// Inicialização do servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
