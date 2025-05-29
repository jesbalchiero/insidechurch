package values

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidEmail = errors.New("email inválido")
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	if !isValidEmail(value) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: strings.ToLower(value)}, nil
}

func (e Email) Value() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return e.value == other.value
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
