package repository

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
	pb "github.com/artem-benda/gophkeeper/client/internal/infrastructure/grpc"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/mapper"
)

// secretRepository implements contract.SecretRepository
type secretRepository struct {
	c pb.GophKeeperServiceClient
}

// NewSecretRepository creates new instance of secretRepository
func NewSecretRepository(c pb.GophKeeperServiceClient) contract.SecretRepository {
	return &secretRepository{c}
}

// Add adds new secret
func (r *secretRepository) Add(c context.Context, name string, encPayload []byte) (string, error) {
	req := &pb.AddSecretRequest{Name: name, Payload: encPayload}
	resp, err := r.c.AddSecret(c, req)
	if err != nil {
		return "", mapper.MapSecretError(err)
	}
	return resp.Guid, nil
}

// Edit edits existing secret
func (r *secretRepository) Edit(ctx context.Context, guid string, name string, encPayload []byte) error {
	req := &pb.UpdateSecretRequest{Guid: guid, Name: name, Payload: encPayload}
	_, err := r.c.UpdateSecret(ctx, req)
	return mapper.MapSecretError(err)
}

// Remove removes existing secret
func (r *secretRepository) Remove(ctx context.Context, guid string) error {
	req := &pb.DeleteSecretRequest{Guid: guid}
	_, err := r.c.DeleteSecret(ctx, req)
	return mapper.MapSecretError(err)
}

// GetAll returns all secrets
func (r *secretRepository) GetAll(ctx context.Context) ([]entity.SecretEncrypted, error) {
	resp, err := r.c.GetAllSecrets(ctx, nil)
	if err != nil {
		return nil, mapper.MapSecretError(err)
	}
	return mapper.MapEncryptedSecrets(resp.Secrets), nil
}

// GetByGUID returns secret by GUID
func (r *secretRepository) GetByGUID(ctx context.Context, guid string) (*entity.SecretEncrypted, error) {
	req := &pb.GetSecretRequest{Guid: guid}
	resp, err := r.c.GetSecret(ctx, req)
	if err != nil {
		return nil, mapper.MapSecretError(err)
	}
	return mapper.MapEncryptedSecret(resp.Secret), nil
}
