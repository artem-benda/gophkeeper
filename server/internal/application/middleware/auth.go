// Package middleware - middleware для GRPC
package middleware

import (
	"context"
)

// DummyAuthFunc - Пустая функция, основная логика проверки авторизации требует наличия проверки в зависимости от метода - см. GophKeeperGrpcServer - AuthFuncOverride
func DummyAuthFunc(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
