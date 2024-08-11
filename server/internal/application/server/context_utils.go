package server

import (
	"context"
)

// getUserIDFromContext returns the user id from the context
func getUserIDFromContext(ctx context.Context) int64 {
	return ctx.Value(UserIDKey).(int64)
}
