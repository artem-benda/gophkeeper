package service

import (
	"context"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"time"
)

var _ contract.SecretService = (*Secret)(nil)

func NewSecretService(repo contract.SecretRepository) contract.SecretService {
	return &Secret{repo: repo}
}

type Secret struct {
	repo contract.SecretRepository
}

func (s Secret) Add(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error) {
	return s.repo.Insert(ctx, userID, guid, name, encPayload, clientTimestamp)
}

func (s Secret) Edit(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error) {
	return s.repo.Update(ctx, userID, guid, name, encPayload, clientTimestamp)
}

func (s Secret) Remove(ctx context.Context, userID int64, guid string) error {
	return s.repo.Delete(ctx, userID, guid)
}

func (s Secret) Get(ctx context.Context, userID int64, guid string) (*entity.Secret, error) {
	return s.repo.Get(ctx, userID, guid)
}

func (s Secret) GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error) {
	return s.repo.GetByUserID(ctx, userID)
}
