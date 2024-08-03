package contract

import (
	"context"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"time"
)

type SecretRepository interface {
	Insert(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error)
	Update(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error)
	Delete(ctx context.Context, userID int64, guid string) error
	Get(ctx context.Context, userID int64, guid string) (*entity.Secret, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error)
}
