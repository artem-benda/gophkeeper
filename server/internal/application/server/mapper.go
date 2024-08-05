package server

import (
	"errors"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/artem-benda/gophkeeper/server/internal/domain/entity"
	pb "github.com/artem-benda/gophkeeper/server/internal/infrastructure/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func mapUserError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, contract.ErrUserAlreadyRegistered):
		return status.Error(codes.AlreadyExists, "user already registered")
	case errors.Is(err, contract.ErrUserNotFound):
		return status.Error(codes.NotFound, "user not found")
	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func mapSecretError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, contract.ErrSecretAlreadyExists):
		return status.Error(codes.AlreadyExists, "secret with provided guid already registered")
	case errors.Is(err, contract.ErrSecretNotFound):
		return status.Error(codes.NotFound, "secret with provided guid not found")
	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func mapToGetSecretResponse(s *entity.Secret) *pb.GetSecretResponse {
	return &pb.GetSecretResponse{
		Secret: &pb.Secret{
			Name:      s.Name,
			Guid:      s.GUID,
			Payload:   s.EncPayload,
			CreatedAt: mapToProtoTimestamp(s.CreatedAt),
			UpdatedAt: mapToProtoTimestamp(s.UpdatedAt),
		},
	}
}

func mapToGetAllSecretsResponse(ss []entity.Secret) *pb.GetAllSecretsResponse {
	secrets := make([]*pb.Secret, len(ss))
	for i, s := range ss {
		secrets[i] = &pb.Secret{
			Name:      s.Name,
			Guid:      s.GUID,
			Payload:   s.EncPayload,
			CreatedAt: mapToProtoTimestamp(s.CreatedAt),
			UpdatedAt: mapToProtoTimestamp(s.UpdatedAt),
		}
	}
	return &pb.GetAllSecretsResponse{Secrets: secrets}
}

func mapToProtoTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
