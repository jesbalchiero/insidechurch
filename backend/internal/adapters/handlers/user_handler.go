package handlers

import (
	userusecase "insidechurch/backend/internal/core/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	getUserUseCase *userusecase.GetUserUseCase
}

func NewUserHandler(getUserUseCase *userusecase.GetUserUseCase) *UserHandler {
	return &UserHandler{getUserUseCase}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	user, err := h.getUserUseCase.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
