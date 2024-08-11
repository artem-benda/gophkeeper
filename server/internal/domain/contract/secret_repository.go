package contract

import (
	"context"
	"time"

	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
)

// SecrettRepository - интерфейс репозитория для работы с секретами
type SecretRepository interface {
	Insert(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error)
	Update(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) error
	Delete(ctx context.Context, userID int64, guid string) error
	Get(ctx context.Context, userID int64, guid string) (*entity.Secret, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error)
}
