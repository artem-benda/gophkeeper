package contract

import (
	"context"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
)

type UserRepository interface {
	Register(ctx context.Context, login string, passwordHash string) (*int64, error)
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
	GetUserByID(ctx context.Context, userID int64) (*entity.User, error)
}
