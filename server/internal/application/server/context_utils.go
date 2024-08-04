package server

import (
	"context"
)

func getUserIDFromContext(ctx context.Context) int64 {
	return ctx.Value(UserIDKey).(int64)
}
