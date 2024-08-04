package middleware

import (
	"context"

	"github.com/artem-benda/gophkeeper/server/internal/application/jwt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var UserIDKey struct{}

func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	userID := jwt.GetUserID(token)
	if userID == -1 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	return context.WithValue(ctx, UserIDKey, userID), nil
}
