package contract

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
)

// SecretRepository - интерфейс для операций с секретами
type SecretRepository interface {
	Add(ctx context.Context, name string, encPayload []byte) (string, error)
	Edit(ctx context.Context, guid string, name string, encPayload []byte) error
	Remove(ctx context.Context, guid string) error
	GetAll(ctx context.Context) ([]entity.SecretEncrypted, error)
	GetByGUID(ctx context.Context, guid string) (*entity.SecretEncrypted, error)
}
