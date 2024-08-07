package service

import (
	"context"
	"time"

	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/google/uuid"
)

var _ contract.SecretService = (*Secret)(nil)

func NewSecretService(repo contract.SecretRepository) contract.SecretService {
	return &Secret{repo: repo}
}

type Secret struct {
	repo contract.SecretRepository
}

func (s Secret) Add(ctx context.Context, userID int64, name string, encPayload []byte) (string, error) {
	guid := uuid.New().String()
	_, err := s.repo.Insert(ctx, userID, guid, name, encPayload, time.Now())
	if err != nil {
		return "", err
	}
	return guid, nil
}

func (s Secret) Edit(ctx context.Context, userID int64, guid string, name string, encPayload []byte) error {
	return s.repo.Update(ctx, userID, guid, name, encPayload, time.Now())
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
