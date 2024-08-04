package server

import (
	"context"

	"github.com/artem-benda/gophkeeper/server/internal/application/middleware"
)

func getUserIDFromContext(ctx context.Context) int64 {
	return ctx.Value(middleware.UserIDKey).(int64)
}
