package contract

import "context"

// UserService - методы бизнес логики для работы с пользователями
type UserService interface {
	Register(ctx context.Context, login string, password string) error
	Login(ctx context.Context, login string, password string) (string, error)
}
