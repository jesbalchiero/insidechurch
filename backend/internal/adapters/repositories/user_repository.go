package repositories

import (
	"errors"
	"insidechurch/backend/internal/core/domain/entities"
	"insidechurch/backend/internal/core/domain/repositories"
	domainerrors "insidechurch/backend/internal/core/errors"

	"gorm.io/gorm"
)

// UserRepository implementa a interface UserRepository usando GORM
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância do UserRepository
func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

// Create implementa a criação de um novo usuário
func (r *UserRepository) Create(user *entities.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return domainerrors.NewInternalError(result.Error)
	}
	return nil
}

// FindByID implementa a busca de usuário por ID
func (r *UserRepository) FindByID(id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.NewUserNotFound(id)
		}
		return nil, domainerrors.NewInternalError(err)
	}

	return &user, nil
}

// FindByEmail implementa a busca de usuário por email
func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.NewUserNotFound(0) // ID não disponível neste caso
		}
		return nil, domainerrors.NewInternalError(err)
	}

	return &user, nil
}

// Update implementa a atualização de um usuário
func (r *UserRepository) Update(user *entities.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return domainerrors.NewInternalError(result.Error)
	}
	return nil
}

// Delete implementa a remoção de um usuário
func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&entities.User{}, id)
	if result.Error != nil {
		return domainerrors.NewInternalError(result.Error)
	}
	if result.RowsAffected == 0 {
		return domainerrors.NewUserNotFound(id)
	}
	return nil
}

// List implementa a listagem de usuários com paginação
func (r *UserRepository) List(page, limit int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	offset := (page - 1) * limit

	// Conta o total de registros
	if err := r.db.Model(&entities.User{}).Count(&total).Error; err != nil {
		return nil, 0, domainerrors.NewInternalError(err)
	}

	// Busca os registros paginados
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, domainerrors.NewInternalError(err)
	}

	return users, total, nil
}

// ExistsByEmail implementa a verificação de existência de email
func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&entities.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, domainerrors.NewInternalError(err)
	}
	return count > 0, nil
}
