package user

import (
	"errors"
	"insidechurch/backend/internal/core/domain/entities"
	"insidechurch/backend/internal/core/ports"
)

type GetUserUseCase struct {
	userRepo ports.UserRepository
}

func NewGetUserUseCase(userRepo ports.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{userRepo}
}

func (uc *GetUserUseCase) GetByID(id uint) (*entities.User, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}
