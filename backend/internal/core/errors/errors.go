package errors

import (
	"fmt"
	"net/http"
)

// DomainError representa um erro do domínio da aplicação
type DomainError struct {
	Code    string
	Message string
	Details map[string]interface{}
	Err     error
}

// Error implementa a interface error
func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap retorna o erro original
func (e *DomainError) Unwrap() error {
	return e.Err
}

// HTTPStatus retorna o código HTTP apropriado para o erro
func (e *DomainError) HTTPStatus() int {
	switch e.Code {
	case ErrUserNotFound:
		return http.StatusNotFound
	case ErrInvalidCredentials:
		return http.StatusUnauthorized
	case ErrEmailAlreadyExists:
		return http.StatusConflict
	case ErrInvalidInput:
		return http.StatusBadRequest
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// Constantes de erro
const (
	ErrUserNotFound       = "USER_NOT_FOUND"
	ErrInvalidCredentials = "INVALID_CREDENTIALS"
	ErrEmailAlreadyExists = "EMAIL_EXISTS"
	ErrInvalidInput       = "INVALID_INPUT"
	ErrUnauthorized       = "UNAUTHORIZED"
	ErrForbidden          = "FORBIDDEN"
	ErrInternal           = "INTERNAL_ERROR"
)

// Funções auxiliares para criar erros do domínio
func NewUserNotFound(userID uint) *DomainError {
	return &DomainError{
		Code:    ErrUserNotFound,
		Message: fmt.Sprintf("Usuário com ID %d não encontrado", userID),
		Details: map[string]interface{}{
			"user_id": userID,
		},
	}
}

func NewInvalidCredentials() *DomainError {
	return &DomainError{
		Code:    ErrInvalidCredentials,
		Message: "Credenciais inválidas",
	}
}

func NewEmailAlreadyExists(email string) *DomainError {
	return &DomainError{
		Code:    ErrEmailAlreadyExists,
		Message: fmt.Sprintf("Email %s já está em uso", email),
		Details: map[string]interface{}{
			"email": email,
		},
	}
}

func NewInvalidInput(message string, details map[string]interface{}) *DomainError {
	return &DomainError{
		Code:    ErrInvalidInput,
		Message: message,
		Details: details,
	}
}

func NewUnauthorized(message string) *DomainError {
	return &DomainError{
		Code:    ErrUnauthorized,
		Message: message,
	}
}

func NewForbidden(message string) *DomainError {
	return &DomainError{
		Code:    ErrForbidden,
		Message: message,
	}
}

func NewInternalError(err error) *DomainError {
	return &DomainError{
		Code:    ErrInternal,
		Message: "Erro interno do servidor",
		Err:     err,
	}
}
