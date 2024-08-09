package ctx

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func WithAuthToken(ctx context.Context, authToken string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", "bearer "+authToken)
}
