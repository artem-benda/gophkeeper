package contract

import (
	"context"

	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
)

// SecretService интерфейс бизнес логики для работы с секретами
type SecretService interface {
	Add(ctx context.Context, userID int64, name string, encPayload []byte) (string, error)
	Edit(ctx context.Context, userID int64, guid string, name string, encPayload []byte) error
	Remove(ctx context.Context, userID int64, guid string) error
	Get(ctx context.Context, userID int64, guid string) (*entity.Secret, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error)
}
