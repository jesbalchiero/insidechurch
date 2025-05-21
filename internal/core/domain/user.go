package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	Password string `json:"-"`
	Name     string `json:"name" validate:"required"`
}
