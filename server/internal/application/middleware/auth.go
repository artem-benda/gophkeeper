package middleware

import (
	"context"
)

// Пустая функция, основная логика проверки авторизации требует наличия проверки в зависимости от метода - см. GophKeeperGrpcServer - AuthFuncOverride
func DummyAuthFunc(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
