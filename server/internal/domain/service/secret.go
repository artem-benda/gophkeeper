package service

import (
	"context"
	"time"

	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/google/uuid"
)

var _ contract.SecretService = (*secret)(nil)

// NewSecretService returns new instance of SecretService
func NewSecretService(repo contract.SecretRepository) contract.SecretService {
	return &secret{repo: repo}
}

// secret is an implementation of contract.SecretService
type secret struct {
	repo contract.SecretRepository
}

// Add adds new secret to database
func (s secret) Add(ctx context.Context, userID int64, name string, encPayload []byte) (string, error) {
	guid := uuid.New().String()
	_, err := s.repo.Insert(ctx, userID, guid, name, encPayload, time.Now())
	if err != nil {
		return "", err
	}
	return guid, nil
}

// Edit edits existing secret in database
func (s secret) Edit(ctx context.Context, userID int64, guid string, name string, encPayload []byte) error {
	return s.repo.Update(ctx, userID, guid, name, encPayload, time.Now())
}

// Remove removes existing secret from database
func (s secret) Remove(ctx context.Context, userID int64, guid string) error {
	return s.repo.Delete(ctx, userID, guid)
}

// Get returns existing secret from database for user and guid
func (s secret) Get(ctx context.Context, userID int64, guid string) (*entity.Secret, error) {
	return s.repo.Get(ctx, userID, guid)
}

// GetByUserID returns existing secrets from database by user ID
func (s secret) GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error) {
	return s.repo.GetByUserID(ctx, userID)
}
