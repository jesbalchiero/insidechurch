package entities

// Role representa um papel no sistema
type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"not null;unique"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

// NewRole cria um novo papel
func NewRole(name string) *Role {
	return &Role{
		Name:        name,
		Permissions: make([]Permission, 0),
	}
}

// AddPermission adiciona uma permissão ao papel
func (r *Role) AddPermission(permission *Permission) {
	for _, p := range r.Permissions {
		if p.Equals(permission) {
			return
		}
	}
	r.Permissions = append(r.Permissions, *permission)
}

// RemovePermission remove uma permissão do papel
func (r *Role) RemovePermission(permission *Permission) {
	for i, p := range r.Permissions {
		if p.Equals(permission) {
			r.Permissions = append(r.Permissions[:i], r.Permissions[i+1:]...)
			return
		}
	}
}

// HasPermission verifica se o papel tem uma determinada permissão
func (r *Role) HasPermission(permission *Permission) bool {
	for _, p := range r.Permissions {
		if p.Equals(permission) {
			return true
		}
	}
	return false
}
