package main

import (
	"os"

	"insidechurch/backend/internal/handlers"
	"insidechurch/backend/internal/middleware"
	"insidechurch/backend/internal/repositories"
	"insidechurch/backend/internal/routes"
	"insidechurch/backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		logrus.Warn("Arquivo .env não encontrado, usando variáveis do sistema.")
	}

	// Configurar logger
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetLevel(logrus.InfoLevel)

	// Configuração do banco de dados usando variáveis de ambiente
	dsn := os.ExpandEnv("host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=${DB_SSLMODE}")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Inicialização dos componentes
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	userHandler := handlers.NewUserHandler(userRepo, authService)
	authMiddleware := middleware.NewAuthMiddleware(authService)
	securityMiddleware := middleware.NewSecurityMiddleware()

	// Configuração do router
	router := gin.Default()

	// Aplicar middlewares de segurança
	router.Use(securityMiddleware.SecurityHeaders())
	router.Use(securityMiddleware.RateLimit())
	router.Use(securityMiddleware.CORS())

	// Configuração das rotas
	routes.SetupRoutes(router, userHandler, authMiddleware)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Porta do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logrus.Infof("Servidor iniciado na porta %s", port)
	if err := router.Run(":" + port); err != nil {
		logrus.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
