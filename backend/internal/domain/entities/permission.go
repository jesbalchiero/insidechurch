package entities

// Permission representa uma permissão no sistema
type Permission struct {
	ID       uint   `gorm:"primaryKey"`
	Resource string `gorm:"not null"`
	Action   string `gorm:"not null"`
}

// NewPermission cria uma nova permissão
func NewPermission(resource, action string) *Permission {
	return &Permission{
		Resource: resource,
		Action:   action,
	}
}

// String retorna a representação em string da permissão
func (p *Permission) String() string {
	return p.Resource + ":" + p.Action
}

// Equals verifica se duas permissões são iguais
func (p *Permission) Equals(other *Permission) bool {
	return p.Resource == other.Resource && p.Action == other.Action
}
