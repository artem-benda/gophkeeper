package repository

import (
	"context"
	"time"

	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	"github.com/artem-benda/gophkeeper/server/internal/infrastructure/db"
)

var _ contract.SecretRepository = (*secretRepository)(nil)

// secretRepository is an implementation of contract.SecretRepository
type secretRepository struct {
	DAO db.SecretDAO
}

// NewSecretRepository returns an implementation of contract.SecretRepository
func NewSecretRepository(dao db.SecretDAO) contract.SecretRepository {
	return &secretRepository{DAO: dao}
}

// Insert inserts new secret
func (s *secretRepository) Insert(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) (*int64, error) {
	return s.DAO.Insert(ctx, userID, guid, name, encPayload, clientTimestamp)
}

// Update updates secret by guid
func (s *secretRepository) Update(ctx context.Context, userID int64, guid string, name string, encPayload []byte, clientTimestamp time.Time) error {
	return s.DAO.Update(ctx, userID, guid, name, encPayload, clientTimestamp)
}

// Delete deletes secret by guid
func (s *secretRepository) Delete(ctx context.Context, userID int64, guid string) error {
	return s.DAO.Delete(ctx, userID, guid)
}

// Get returns secret by guid and user id
func (s *secretRepository) Get(ctx context.Context, userID int64, guid string) (*entity.Secret, error) {
	return s.DAO.GetByGUID(ctx, userID, guid)
}

// GetByGUID returns secrets by user id
func (s *secretRepository) GetByUserID(ctx context.Context, userID int64) ([]entity.Secret, error) {
	return s.DAO.GetByUserID(ctx, userID)
}
