package mapper

import (
	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/client/internal/domain/entity"
	pb "github.com/artem-benda/gophkeeper/client/internal/infrastructure/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapSecretError(grpcError error) error {
	if grpcError == nil {
		return nil
	}
	if e, ok := status.FromError(grpcError); ok {
		switch e.Code() {
		case codes.AlreadyExists:
			return contract.ErrSecretAlreadyExists
		case codes.InvalidArgument:
			return contract.ErrSecretNotFound
		default:
			return grpcError
		}
	} else {
		return grpcError
	}
}

func MapEncryptedSecret(s *pb.Secret) *entity.SecretEncrypted {
	if s == nil {
		return nil
	}
	return &entity.SecretEncrypted{
		GUID:       s.Guid,
		Name:       s.Name,
		CreatedAt:  s.CreatedAt.AsTime(),
		UpdatedAt:  s.UpdatedAt.AsTime(),
		EncPayload: s.Payload,
	}
}

func MapEncryptedSecrets(s []*pb.Secret) []entity.SecretEncrypted {
	secrets := make([]entity.SecretEncrypted, len(s))
	for ind, v := range s {
		secrets[ind] = *MapEncryptedSecret(v)
	}
	return secrets
}
