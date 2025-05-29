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

	// Grupo de rotas do Swagger
	swagger := router.Group("/api")
	swagger.Use(securityMiddleware.SecurityHeaders())
	swagger.Use(securityMiddleware.RateLimit())
	swagger.Use(securityMiddleware.CORS())
	{
		// Servir arquivos estáticos do Swagger
		swagger.Static("/swagger-static", "/app/swagger")

		// Configuração do Swagger UI
		url := ginSwagger.URL("/api/swagger-static/swagger.json")
		swagger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	// Grupo de rotas públicas
	public := router.Group("/api")
	public.Use(securityMiddleware.SecurityHeaders())
	public.Use(securityMiddleware.RateLimit())
	public.Use(securityMiddleware.CORS())
	{
		// Grupo de rotas de autenticação
		auth := public.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	// Grupo de rotas protegidas
	protected := router.Group("/api")
	protected.Use(securityMiddleware.SecurityHeaders())
	protected.Use(securityMiddleware.RateLimit())
	protected.Use(securityMiddleware.CORS())
	protected.Use(authMiddleware.Authenticate())
	{
		protected.GET("/user", userHandler.GetUser)
	}
}
