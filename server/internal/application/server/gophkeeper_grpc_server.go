package server

import (
	"context"
	pb "github.com/artem-benda/gophkeeper/server/internal/application/grpc"
	"github.com/artem-benda/gophkeeper/server/internal/application/jwt"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GophKeeperGrpcServer struct {
	pb.UnimplementedGophKeeperServiceServer
	USvc contract.UserService
	SSvc contract.SecretService
}

func (s *GophKeeperGrpcServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.Login == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login or password cannot be empty")
	}

	userID, err := s.USvc.Register(ctx, req.Login, req.Password)
	if err != nil {
		return nil, mapUserError(err)
	}

	token, err := jwt.BuildJWTString(*userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "build jwt error")
	}

	return &pb.RegisterResponse{Token: token}, nil
}

func (s *GophKeeperGrpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Login == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login or password cannot be empty")
	}

	userID, err := s.USvc.Login(ctx, req.Login, req.Password)
	if err != nil {
		return nil, mapUserError(err)
	}

	token, err := jwt.BuildJWTString(*userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "build jwt error")
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (s *GophKeeperGrpcServer) GetSecret(ctx context.Context, req *pb.GetSecretRequest) (*pb.GetSecretResponse, error) {
	userID := getUserIDFromContext(ctx)
	secret, err := s.SSvc.Get(ctx, userID, req.GetGuid())
	if err != nil {
		return nil, mapSecretError(err)
	}

	return mapToGetSecretResponse(secret), nil
}

func (s *GophKeeperGrpcServer) GetAllSecrets(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllSecretsResponse, error) {
	userID := getUserIDFromContext(ctx)
	secrets, err := s.SSvc.GetByUserID(ctx, userID)
	if err != nil {
		return nil, mapSecretError(err)
	}

	return mapToGetAllSecretsResponse(secrets), nil
}

func (s *GophKeeperGrpcServer) AddSecret(ctx context.Context, req *pb.AddSecretRequest) (*emptypb.Empty, error) {
	userID := getUserIDFromContext(ctx)
	_, err := s.SSvc.Add(ctx, userID, req.Guid, req.Name, req.Payload, req.ClientTimestamp.AsTime())
	if err != nil {
		return nil, mapSecretError(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *GophKeeperGrpcServer) UpdateSecret(ctx context.Context, req *pb.UpdateSecretRequest) (*emptypb.Empty, error) {
	userID := getUserIDFromContext(ctx)
	_, err := s.SSvc.Edit(ctx, userID, req.Guid, req.Name, req.Payload, req.ClientTimestamp.AsTime())
	if err != nil {
		return nil, mapSecretError(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *GophKeeperGrpcServer) DeleteSecret(ctx context.Context, req *pb.DeleteSecretRequest) (*emptypb.Empty, error) {
	userID := getUserIDFromContext(ctx)
	err := s.SSvc.Remove(ctx, userID, req.Guid)
	if err != nil {
		return nil, mapSecretError(err)
	}

	return &emptypb.Empty{}, nil
}
