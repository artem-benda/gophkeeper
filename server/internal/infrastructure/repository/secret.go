package repository

import (
	"context"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/artem-benda/gophkeeper/server/internal/infrastructure/db"
	"time"
)

var _ contract.SecretRepository = (*SecretRepository)(nil)

type SecretRepository struct {
	DAO db.SecretDAO
}

func (s SecretRepository) Insert(ctx context.Context, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error) {
	return s.Insert(ctx, guid, name, encPayload, clientTimestamp)
}

func (s SecretRepository) Update(ctx context.Context, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error) {
	return s.Update(ctx, guid, name, encPayload, clientTimestamp)
}

func (s SecretRepository) Delete(ctx context.Context, guid string) error {
	return s.Delete(ctx, guid)
}

func (s SecretRepository) Get(ctx context.Context, guid string) (*entity.Secret, error) {
	return s.Get(ctx, guid)
}

func (s SecretRepository) GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error) {
	return s.GetByUserID(ctx, userID)
}
