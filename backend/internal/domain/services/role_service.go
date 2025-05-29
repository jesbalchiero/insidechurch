package services

import (
	"errors"

	"insidechurch/backend/internal/domain/entities"
	"insidechurch/backend/internal/domain/repositories"
)

type RoleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) *RoleService {
	return &RoleService{
		roleRepo: roleRepo,
	}
}

// CreateRole cria um novo papel
func (s *RoleService) CreateRole(name string) (*entities.Role, error) {
	existingRole, _ := s.roleRepo.FindByName(name)
	if existingRole != nil {
		return nil, errors.New("papel já existe")
	}

	role := entities.NewRole(name)
	err := s.roleRepo.Create(role)
	if err != nil {
		return nil, err
	}

	return role, nil
}

// AssignPermission atribui uma permissão a um papel
func (s *RoleService) AssignPermission(roleID uint, resource, action string) error {
	_, err := s.roleRepo.FindByID(roleID)
	if err != nil {
		return err
	}

	permission := entities.NewPermission(resource, action)
	return s.roleRepo.AddPermission(roleID, permission)
}

// RemovePermission remove uma permissão de um papel
func (s *RoleService) RemovePermission(roleID uint, resource, action string) error {
	_, err := s.roleRepo.FindByID(roleID)
	if err != nil {
		return err
	}

	permission := entities.NewPermission(resource, action)
	return s.roleRepo.RemovePermission(roleID, permission)
}

// CheckPermission verifica se um papel tem uma determinada permissão
func (s *RoleService) CheckPermission(roleID uint, resource, action string) (bool, error) {
	role, err := s.roleRepo.FindByID(roleID)
	if err != nil {
		return false, err
	}

	permission := entities.NewPermission(resource, action)
	return role.HasPermission(permission), nil
}

// ListRoles retorna todos os papéis
func (s *RoleService) ListRoles() ([]entities.Role, error) {
	return s.roleRepo.List()
}

// GetRole retorna um papel pelo ID
func (s *RoleService) GetRole(id uint) (*entities.Role, error) {
	return s.roleRepo.FindByID(id)
}

// DeleteRole remove um papel
func (s *RoleService) DeleteRole(id uint) error {
	return s.roleRepo.Delete(id)
}
