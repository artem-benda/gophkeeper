package ctx

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// WithAuthToken добавить токен авторизации в grpc контекст
func WithAuthToken(ctx context.Context, authToken string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", "bearer "+authToken)
}
