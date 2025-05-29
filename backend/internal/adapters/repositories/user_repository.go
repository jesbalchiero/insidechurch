package repositories

import (
	"insidechurch/backend/internal/core/domain/entities"
	"insidechurch/backend/internal/core/ports"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância do repositório de usuários
func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create cria um novo usuário
func (r *userRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

// FindByID busca um usuário pelo ID
func (r *userRepository) FindByID(id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail busca um usuário pelo email
func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update atualiza um usuário
func (r *userRepository) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

// Delete remove um usuário
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entities.User{}, id).Error
}
