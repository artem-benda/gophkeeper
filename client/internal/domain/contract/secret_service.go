package contract

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
)

type SecretService interface {
	Add(ctx context.Context, name string, encPayload []byte) (string, error)
	Edit(ctx context.Context, guid string, name string, encPayload []byte) error
	Remove(ctx context.Context, guid string) error
	GetAll(ctx context.Context) ([]entity.Secret, error)
	GetByGUID(ctx context.Context, guid string) (*entity.Secret, error)
}
