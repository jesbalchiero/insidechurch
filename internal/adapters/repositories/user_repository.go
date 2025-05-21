package repositories

import (
	"insidechurch/internal/core/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id uint) (*domain.User, error)
}

type PostgresUserRepository struct {
	DB *gorm.DB
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	// Implementação de criação
	return r.DB.Create(user).Error
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	// Implementação de busca por email
	var user domain.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) FindByID(id uint) (*domain.User, error) {
	// Implementação de busca por ID
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
