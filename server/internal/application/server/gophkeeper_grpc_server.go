// Package server - реализация GRPC сервера
package server

import (
	"context"
	"log/slog"

	"github.com/artem-benda/gophkeeper/server/internal/application/jwt"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	pb "github.com/artem-benda/gophkeeper/server/internal/infrastructure/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GophKeeperGrpcServer - реализация GRPC сервера
type GophKeeperGrpcServer struct {
	pb.UnimplementedGophKeeperServiceServer
	USvc contract.UserService
	SSvc contract.SecretService
}

// Register - Регистрация пользователя
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

// Login - Аутентификация пользователя
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

// GetSecret - Получить запись секретной информации по guid для пользователя
func (s *GophKeeperGrpcServer) GetSecret(ctx context.Context, req *pb.GetSecretRequest) (*pb.GetSecretResponse, error) {
	userID := getUserIDFromContext(ctx)
	secret, err := s.SSvc.Get(ctx, userID, req.GetGuid())
	if err != nil {
		return nil, mapSecretError(err)
	}

	return mapToGetSecretResponse(secret), nil
}

// GetAllSecrets - Получить все записи секретной информации для пользователя
func (s *GophKeeperGrpcServer) GetAllSecrets(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllSecretsResponse, error) {
	userID := getUserIDFromContext(ctx)
	secrets, err := s.SSvc.GetByUserID(ctx, userID)
	if err != nil {
		return nil, mapSecretError(err)
	}

	return mapToGetAllSecretsResponse(secrets), nil
}

// AddSecret - Добавить запись секретной информации
func (s *GophKeeperGrpcServer) AddSecret(ctx context.Context, req *pb.AddSecretRequest) (*emptypb.Empty, error) {
	userID := getUserIDFromContext(ctx)
	_, err := s.SSvc.Add(ctx, userID, req.Guid, req.Name, req.Payload, req.ClientTimestamp.AsTime())
	if err != nil {
		return nil, mapSecretError(err)
	}

	return &emptypb.Empty{}, nil
}

// UpdateSecret - Изменить запись секретной информации
func (s *GophKeeperGrpcServer) UpdateSecret(ctx context.Context, req *pb.UpdateSecretRequest) (*emptypb.Empty, error) {
	userID := getUserIDFromContext(ctx)
	err := s.SSvc.Edit(ctx, userID, req.Guid, req.Name, req.Payload, req.ClientTimestamp.AsTime())
	if err != nil {
		return nil, mapSecretError(err)
	}

	return &emptypb.Empty{}, nil
}

// DeleteSecret - Удалить запись секретной информации
func (s *GophKeeperGrpcServer) DeleteSecret(ctx context.Context, req *pb.DeleteSecretRequest) (*emptypb.Empty, error) {
	userID := getUserIDFromContext(ctx)
	err := s.SSvc.Remove(ctx, userID, req.Guid)
	if err != nil {
		return nil, mapSecretError(err)
	}

	return &emptypb.Empty{}, nil
}

var UserIDKey struct{}

// AuthFuncOverride - Метод, используемый middleware Auth - имеет приоритет над AuthFunc
func (s *GophKeeperGrpcServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	slog.Debug("client is calling method: ", fullMethodName)
	// Для этих методов не нужна авторизация
	if fullMethodName == pb.GophKeeperService_Register_FullMethodName || fullMethodName == pb.GophKeeperService_Login_FullMethodName {
		return ctx, nil
	}

	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	userID := jwt.GetUserID(token)
	if userID == -1 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	return context.WithValue(ctx, UserIDKey, userID), nil
}
