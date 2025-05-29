package integration

import (
	"testing"

	"insidechurch/backend/internal/adapters/repositories"
	"insidechurch/backend/internal/core/domain/entities"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	// Configurar banco de dados de teste
	dsn := "host=localhost user=postgres password=postgres dbname=insidechurch_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados de teste: %v", err)
	}

	// Limpar tabelas antes dos testes
	db.Exec("DROP TABLE IF EXISTS users CASCADE")
	db.AutoMigrate(&entities.User{})

	return db
}

func TestUserRepository(t *testing.T) {
	db := setupTestDB(t)
	userRepo := repositories.NewUserRepository(db)

	// Teste de criação de usuário
	t.Run("Criar usuário", func(t *testing.T) {
		user := &entities.User{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "Test@123",
		}

		err := userRepo.Create(user)
		assert.NoError(t, err)
		assert.NotZero(t, user.ID)
	})

	// Teste de busca de usuário por email
	t.Run("Buscar usuário por email", func(t *testing.T) {
		user, err := userRepo.FindByEmail("test@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "Test User", user.Name)
		assert.Equal(t, "test@example.com", user.Email)
	})

	// Teste de busca de usuário por ID
	t.Run("Buscar usuário por ID", func(t *testing.T) {
		// Primeiro criar um usuário
		user := &entities.User{
			Name:     "Another User",
			Email:    "another@example.com",
			Password: "Test@123",
		}
		err := userRepo.Create(user)
		assert.NoError(t, err)

		// Depois buscar pelo ID
		foundUser, err := userRepo.FindByID(user.ID)
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.Name, foundUser.Name)
		assert.Equal(t, user.Email, foundUser.Email)
	})

	// Teste de atualização de usuário
	t.Run("Atualizar usuário", func(t *testing.T) {
		// Primeiro criar um usuário
		user := &entities.User{
			Name:     "Update User",
			Email:    "update@example.com",
			Password: "Test@123",
		}
		err := userRepo.Create(user)
		assert.NoError(t, err)

		// Atualizar o usuário
		user.Name = "Updated Name"
		err = userRepo.Update(user)
		assert.NoError(t, err)

		// Verificar se foi atualizado
		updatedUser, err := userRepo.FindByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updatedUser.Name)
	})

	// Teste de exclusão de usuário
	t.Run("Excluir usuário", func(t *testing.T) {
		// Primeiro criar um usuário
		user := &entities.User{
			Name:     "Delete User",
			Email:    "delete@example.com",
			Password: "Test@123",
		}
		err := userRepo.Create(user)
		assert.NoError(t, err)

		// Excluir o usuário
		err = userRepo.Delete(user.ID)
		assert.NoError(t, err)

		// Verificar se foi excluído
		_, err = userRepo.FindByID(user.ID)
		assert.Error(t, err)
	})
}
