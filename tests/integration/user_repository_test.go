package integration

import (
	"testing"

	"insidechurch/internal/core/domain"
	"insidechurch/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	// Configurar banco de dados de teste
	db := setupTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Dados de teste
	user := &domain.User{
		Email:    "teste@exemplo.com",
		Name:     "Usuário Teste",
		Password: "senha123",
	}

	// Teste de criação
	t.Run("Criar usuário", func(t *testing.T) {
		err := repo.Create(user)
		assert.NoError(t, err)
		assert.NotZero(t, user.ID)
	})

	// Teste de busca por email
	t.Run("Buscar usuário por email", func(t *testing.T) {
		foundUser, err := repo.FindByEmail(user.Email)
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user.Email, foundUser.Email)
		assert.Equal(t, user.Name, foundUser.Name)
	})

	// Teste de busca por ID
	t.Run("Buscar usuário por ID", func(t *testing.T) {
		foundUser, err := repo.FindByID(user.ID)
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user.Email, foundUser.Email)
		assert.Equal(t, user.Name, foundUser.Name)
	})

	// Teste de atualização
	t.Run("Atualizar usuário", func(t *testing.T) {
		user.Name = "Nome Atualizado"
		err := repo.Update(user)
		assert.NoError(t, err)

		// Verificar se foi atualizado
		foundUser, err := repo.FindByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Nome Atualizado", foundUser.Name)
	})

	// Teste de exclusão
	t.Run("Excluir usuário", func(t *testing.T) {
		err := repo.Delete(user.ID)
		assert.NoError(t, err)

		// Verificar se foi excluído
		foundUser, err := repo.FindByID(user.ID)
		assert.NoError(t, err)
		assert.Nil(t, foundUser)
	})

	// Teste de usuário não encontrado
	t.Run("Buscar usuário inexistente", func(t *testing.T) {
		foundUser, err := repo.FindByEmail("naoexiste@exemplo.com")
		assert.NoError(t, err)
		assert.Nil(t, foundUser)
	})
}
