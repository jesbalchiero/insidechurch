package routes

import (
	"insidechurch/internal/handlers"
	"insidechurch/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Grupo de rotas públicas
	public := router.Group("/api")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	}

	// Grupo de rotas protegidas
	protected := router.Group("/api")
	protected.Use(authMiddleware.Authenticate())
	{
		protected.GET("/user", userHandler.GetUser)
	}
}
