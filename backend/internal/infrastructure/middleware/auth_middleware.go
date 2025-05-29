package middleware

import (
	"insidechurch/backend/internal/domain/services"
	"net/http"
)

type AuthMiddleware struct {
	roleService *services.RoleService
}

func NewAuthMiddleware(roleService *services.RoleService) *AuthMiddleware {
	return &AuthMiddleware{
		roleService: roleService,
	}
}

// RequirePermission cria um middleware que verifica se o usuário tem a permissão necessária
func (m *AuthMiddleware) RequirePermission(resource, action string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Obter o roleID do usuário autenticado
			// Por enquanto, vamos usar um roleID fixo para teste
			roleID := uint(1)

			hasPermission, err := m.roleService.CheckPermission(roleID, resource, action)
			if err != nil || !hasPermission {
				http.Error(w, "Acesso negado", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAnyPermission cria um middleware que verifica se o usuário tem pelo menos uma das permissões necessárias
func (m *AuthMiddleware) RequireAnyPermission(permissions []struct{ Resource, Action string }) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Obter o roleID do usuário autenticado
			// Por enquanto, vamos usar um roleID fixo para teste
			roleID := uint(1)

			for _, perm := range permissions {
				hasPermission, err := m.roleService.CheckPermission(roleID, perm.Resource, perm.Action)
				if err == nil && hasPermission {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Acesso negado", http.StatusForbidden)
		})
	}
}

// RequireAllPermissions cria um middleware que verifica se o usuário tem todas as permissões necessárias
func (m *AuthMiddleware) RequireAllPermissions(permissions []struct{ Resource, Action string }) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Obter o roleID do usuário autenticado
			// Por enquanto, vamos usar um roleID fixo para teste
			roleID := uint(1)

			for _, perm := range permissions {
				hasPermission, err := m.roleService.CheckPermission(roleID, perm.Resource, perm.Action)
				if err != nil || !hasPermission {
					http.Error(w, "Acesso negado", http.StatusForbidden)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
