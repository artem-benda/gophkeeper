package contract

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
)

// SecretRepository - интерфейс для бизнес логики работы с секретами
type SecretService interface {
	AddLoginPassword(ctx context.Context, name string, login string, password string, metadata string) (string, error)
	AddText(ctx context.Context, name string, text string, metadata string) (string, error)
	AddBinary(ctx context.Context, name string, data []byte, metadata string) (string, error)
	AddBankingCard(ctx context.Context, name string, number string, owner string, dueTo string, cvv string, metadata string) (string, error)
	ReplaceWithLoginPassword(ctx context.Context, guid string, name string, login string, password string, metadata string) error
	ReplaceWithText(ctx context.Context, guid string, name string, text string, metadata string) error
	ReplaceWithBinary(ctx context.Context, guid string, name string, data []byte, metadata string) error
	ReplaceWithBankingCard(ctx context.Context, guid string, name string, number string, owner string, dueTo string, cvv string, metadata string) error
	Remove(ctx context.Context, guid string) error
	GetAll(ctx context.Context) ([]entity.Secret, error)
	GetByGUID(ctx context.Context, guid string) (*entity.Secret, error)
}
