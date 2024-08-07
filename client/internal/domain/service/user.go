package service

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
)

// user - реаализация интерфейса contract.UserService
type user struct {
	r contract.UserRepository
}

// NewUserService - создать экземпляр с интерфейсом contract.UserService
func NewUserService(r contract.UserRepository) contract.UserService {
	return &user{r}
}

// Login - логин пользователя
func (s *user) Login(ctx context.Context, login string, password string) (string, error) {
	return s.r.Login(ctx, login, password)
}

// Register - регистрация пользователя
func (s *user) Register(ctx context.Context, login string, password string) error {
	return s.r.Register(ctx, login, password)
}
