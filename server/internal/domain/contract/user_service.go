package contract

import (
	"context"
)

type UserService interface {
	Register(ctx context.Context, login string, passwordHash string) (*int64, error)
	Login(ctx context.Context, login string, password string) (*int64, error)
}
