package repositories

import "insidechurch/backend/internal/domain/entities"

// RoleRepository define as operações de persistência para roles
type RoleRepository interface {
	Create(role *entities.Role) error
	FindByID(id uint) (*entities.Role, error)
	FindByName(name string) (*entities.Role, error)
	Update(role *entities.Role) error
	Delete(id uint) error
	List() ([]entities.Role, error)
	AddPermission(roleID uint, permission *entities.Permission) error
	RemovePermission(roleID uint, permission *entities.Permission) error
}
