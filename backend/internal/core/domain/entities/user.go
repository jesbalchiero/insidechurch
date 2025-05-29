package entities

import (
	"time"

	"gorm.io/gorm"
)

// User representa a entidade de usuário no domínio
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"` // "-" para não serializar a senha
}

// TableName especifica o nome da tabela no banco de dados
func (User) TableName() string {
	return "users"
}

// BeforeCreate é um hook do GORM que é executado antes de criar um usuário
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate é um hook do GORM que é executado antes de atualizar um usuário
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
