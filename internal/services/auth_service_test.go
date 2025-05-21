package services

import (
	"insidechurch/internal/core/domain"
	"testing"

	"gorm.io/gorm"
)

type mockUserRepo struct {
	user    *domain.User
	findErr error
}

func (m *mockUserRepo) FindByEmail(email string) (*domain.User, error) {
	return m.user, m.findErr
}
func (m *mockUserRepo) Create(user *domain.User) error         { return nil }
func (m *mockUserRepo) FindByID(id uint) (*domain.User, error) { return nil, nil }
func (m *mockUserRepo) Update(user *domain.User) error         { return nil }
func (m *mockUserRepo) Delete(id uint) error                   { return nil }

func TestValidatePassword(t *testing.T) {
	s := NewAuthService(nil)
	if err := s.ValidatePassword("123"); err == nil {
		t.Error("deveria falhar para senha curta")
	}
	if err := s.ValidatePassword("123456"); err != nil {
		t.Error("não deveria falhar para senha válida")
	}
}

func TestHashAndCheckPassword(t *testing.T) {
	s := NewAuthService(nil)
	hash, err := s.HashPassword("senha123")
	if err != nil {
		t.Fatal(err)
	}
	if !s.CheckPasswordHash("senha123", hash) {
		t.Error("hash não confere com senha")
	}
	if s.CheckPasswordHash("errada", hash) {
		t.Error("hash não deveria conferir com senha errada")
	}
}

func TestGenerateAndValidateToken(t *testing.T) {
	s := NewAuthService(nil)
	user := &domain.User{Model: gorm.Model{ID: 1}, Email: "a@b.com", Name: "Teste"}
	token, err := s.GenerateToken(user)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := s.ValidateToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if claims.UserID != user.ID {
		t.Errorf("esperado userID %d, obtido %d", user.ID, claims.UserID)
	}
}

func TestAuthenticate(t *testing.T) {
	user := &domain.User{Model: gorm.Model{ID: 1}, Email: "a@b.com", Name: "Teste"}
	s := NewAuthService(&mockUserRepo{user: user})
	hash, _ := s.HashPassword("senha123")
	user.Password = hash

	token, err := s.Authenticate("a@b.com", "senha123")
	if err != nil || token == "" {
		t.Error("autenticação deveria funcionar com senha correta")
	}

	_, err = s.Authenticate("a@b.com", "errada")
	if err == nil {
		t.Error("autenticação deveria falhar com senha errada")
	}

	s2 := NewAuthService(&mockUserRepo{user: nil})
	_, err = s2.Authenticate("naoexiste@b.com", "senha123")
	if err == nil {
		t.Error("autenticação deveria falhar para usuário inexistente")
	}
}
