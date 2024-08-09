package repository

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
	pb "github.com/artem-benda/gophkeeper/client/internal/infrastructure/grpc"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/mapper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type secretRepository struct {
	c pb.GophKeeperServiceClient
}

func NewSecretRepository(c pb.GophKeeperServiceClient) contract.SecretRepository {
	return &secretRepository{c}
}

func (r *secretRepository) Register(ctx context.Context, login string, password string) error {
	req := &pb.RegisterRequest{Login: login, Password: password}
	_, err := r.c.Register(ctx, req)
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.AlreadyExists:
			return contract.ErrUserAlreadyExists
		case codes.InvalidArgument:
			return contract.ErrUserInvalidCredentials
		default:
			return err
		}
	}
	return err
}

func (r *secretRepository) Add(ctx context.Context, name string, encPayload []byte) (string, error) {
	req := &pb.AddSecretRequest{Name: name, Payload: encPayload}
	resp, err := r.c.AddSecret(ctx, req)
	if err != nil {
		return "", mapper.MapSecretError(err)
	}
	return resp.Guid, nil
}

func (r *secretRepository) Edit(ctx context.Context, guid string, name string, encPayload []byte) error {
	req := &pb.UpdateSecretRequest{Guid: guid, Name: name, Payload: encPayload}
	_, err := r.c.UpdateSecret(ctx, req)
	return mapper.MapSecretError(err)
}

func (r *secretRepository) Remove(ctx context.Context, guid string) error {
	req := &pb.DeleteSecretRequest{Guid: guid}
	_, err := r.c.DeleteSecret(ctx, req)
	return mapper.MapSecretError(err)
}

func (r *secretRepository) GetAll(ctx context.Context) ([]entity.SecretEncrypted, error) {
	resp, err := r.c.GetAllSecrets(ctx, nil)
	if err != nil {
		return nil, mapper.MapSecretError(err)
	}
	return mapper.MapEncryptedSecrets(resp.Secrets), nil
}

func (r *secretRepository) GetByGUID(ctx context.Context, guid string) (*entity.SecretEncrypted, error) {
	req := &pb.GetSecretRequest{Guid: guid}
	resp, err := r.c.GetSecret(ctx, req)
	if err != nil {
		return nil, mapper.MapSecretError(err)
	}
	return mapper.MapEncryptedSecret(resp.Secret), nil
}
