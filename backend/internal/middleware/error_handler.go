package middleware

import (
	"insidechurch/backend/internal/core/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse representa a estrutura padrão de resposta de erro
type ErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ErrorHandler é um middleware que trata erros de forma consistente
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Verifica se há erros
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Tenta converter para DomainError
			if domainErr, ok := err.(*errors.DomainError); ok {
				c.JSON(domainErr.HTTPStatus(), ErrorResponse{
					Code:    domainErr.Code,
					Message: domainErr.Message,
					Details: domainErr.Details,
				})
				return
			}

			// Erro genérico
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    errors.ErrInternal,
				Message: "Erro interno do servidor",
			})
		}
	}
}
