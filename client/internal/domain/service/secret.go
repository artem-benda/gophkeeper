package service

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
)

type secret struct {
	r contract.SecretRepository
}

func NewSecretService(r contract.SecretRepository) contract.SecretService {
	return &secret{r}
}

func (s *secret) Add(ctx context.Context, name string, encPayload []byte) (string, error) {
	return s.r.Add(ctx, name, encPayload)
}

func (s *secret) Edit(ctx context.Context, guid string, name string, encPayload []byte) error {
	return s.r.Edit(ctx, guid, name, encPayload)
}

func (s *secret) Remove(ctx context.Context, guid string) error {
	return s.r.Remove(ctx, guid)
}

func (s *secret) GetAll(ctx context.Context) ([]entity.Secret, error) {
	return s.r.GetAll(ctx)
}

func (s *secret) GetByGUID(ctx context.Context, guid string) (*entity.Secret, error) {
	return s.r.GetByGUID(ctx, guid)
}
