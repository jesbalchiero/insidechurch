package main

import (
	"log"
	"os"

	_ "insidechurch/backend/cmd/api/docs" // Importar a documentação do Swagger
	"insidechurch/backend/internal/adapters/handlers"
	"insidechurch/backend/internal/adapters/repositories"
	"insidechurch/backend/internal/core/usecases/auth"
	"insidechurch/backend/internal/core/usecases/user"
	"insidechurch/backend/internal/middleware"
	"insidechurch/backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Inside Church API
// @version         1.0
// @description     API para o sistema Inside Church
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado")
	}

	// Configura o modo do Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
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

	// Inicializa o router
	router := gin.Default()

	// Inicializa os repositórios
	userRepo := repositories.NewUserRepository(db)

	// Inicializa os casos de uso
	loginUseCase := auth.NewLoginUseCase(userRepo)
	registerUseCase := auth.NewRegisterUseCase(userRepo)
	getUserUseCase := user.NewGetUserUseCase(userRepo)

	// Inicializa os middlewares
	authMiddleware := middleware.NewAuthMiddleware(loginUseCase)
	securityMiddleware := middleware.NewSecurityMiddleware()

	// Inicializa os handlers
	authHandler := handlers.NewAuthHandler(loginUseCase, registerUseCase)
	userHandler := handlers.NewUserHandler(getUserUseCase)

	// Configura as rotas
	routes.SetupRoutes(router, authHandler, userHandler, authMiddleware, securityMiddleware)

	// Inicia o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logrus.Infof("Servidor iniciado na porta %s", port)
	if err := router.Run(":" + port); err != nil {
		logrus.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
