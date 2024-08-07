package contract

import "context"

// UserRepository - операции с пользователями
type UserRepository interface {
	Login(ctx context.Context, username string, password string) (string, error)
	Register(ctx context.Context, login string, password string) error
}
