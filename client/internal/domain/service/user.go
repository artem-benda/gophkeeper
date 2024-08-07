package service

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
)

type user struct {
	r contract.UserRepository
}

func NewUserService(r contract.UserRepository) contract.UserService {
	return &user{r}
}

func (s *user) Login(ctx context.Context, login string, password string) (string, error) {
	return s.r.Login(ctx, login, password)
}

func (s *user) Register(ctx context.Context, login string, password string) error {
	return s.r.Register(ctx, login, password)
}
