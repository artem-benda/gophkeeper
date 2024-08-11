package repository

import (
	"context"

	"github.com/artem-benda/gophkeeper/client/internal/domain/contract"
	pb "github.com/artem-benda/gophkeeper/client/internal/infrastructure/grpc"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/mapper"
)

// userRepository is implementation of contract.UserRepository
type userRepository struct {
	c pb.GophKeeperServiceClient
}

// NewUserRepository returns implementation of contract.UserRepository
func NewUserRepository(c pb.GophKeeperServiceClient) contract.UserRepository {
	return &userRepository{c}
}

// Register registers new user
func (r *userRepository) Register(ctx context.Context, login string, password string) error {
	req := &pb.RegisterRequest{Login: login, Password: password}
	_, err := r.c.Register(ctx, req)
	return mapper.MapUserError(err)
}

// Login tries to log in user
func (r *userRepository) Login(ctx context.Context, login string, password string) (string, error) {
	req := &pb.LoginRequest{Login: login, Password: password}
	resp, err := r.c.Login(ctx, req)
	if err != nil {
		return "", mapper.MapUserError(err)
	}
	return resp.Token, nil
}
