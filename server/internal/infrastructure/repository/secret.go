package repository

import (
	"context"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/artem-benda/gophkeeper/server/internal/infrastructure/db"
	"time"
)

var _ contract.SecretRepository = (*secretRepository)(nil)

type secretRepository struct {
	DAO db.SecretDAO
}

func NewSecretRepository(dao db.SecretDAO) contract.SecretRepository {
	return &secretRepository{DAO: dao}
}

func (s *secretRepository) Insert(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error) {
	return s.Insert(ctx, userID, guid, name, encPayload, clientTimestamp)
}

func (s *secretRepository) Update(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error) {
	return s.Update(ctx, userID, guid, name, encPayload, clientTimestamp)
}

func (s *secretRepository) Delete(ctx context.Context, userID int64, guid string) error {
	return s.Delete(ctx, userID, guid)
}

func (s *secretRepository) Get(ctx context.Context, userID int64, guid string) (*entity.Secret, error) {
	return s.Get(ctx, userID, guid)
}

func (s *secretRepository) GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error) {
	return s.GetByUserID(ctx, userID)
}
