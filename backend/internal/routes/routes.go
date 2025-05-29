package routes

import (
	"insidechurch/backend/internal/adapters/handlers"
	"insidechurch/backend/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(
	router *gin.Engine,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	authMiddleware *middleware.AuthMiddleware,
	securityMiddleware *middleware.SecurityMiddleware,
) {
	// Middlewares globais
	router.Use(securityMiddleware.SecurityHeaders())
	router.Use(securityMiddleware.RateLimit())
	router.Use(securityMiddleware.CORS())
	router.Use(middleware.ErrorHandler())

	// Rota raiz
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/api/swagger/index.html")
	})

	// Rota de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API is running",
		})
	})

	// Grupo de rotas da API
	api := router.Group("/api")
	{
		// Grupo de rotas do Swagger
		swagger := api.Group("/swagger")
		{
			// Servir arquivos estáticos do Swagger
			swagger.Static("/static", "/app/swagger")

			// Configuração do Swagger UI
			url := ginSwagger.URL("/api/swagger/static/swagger.json")
			swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		}

		// Grupo de rotas públicas
		public := api.Group("")
		{
			// Rotas de autenticação
			auth := public.Group("/auth")
			{
				auth.POST("/register", authHandler.Register)
				auth.POST("/login", authHandler.Login)
			}
		}

		// Grupo de rotas protegidas
		protected := api.Group("")
		protected.Use(authMiddleware.Authenticate())
		{
			// Rotas de usuário
			users := protected.Group("/users")
			{
				users.GET("/me", userHandler.GetUser)
			}
		}
	}
}
