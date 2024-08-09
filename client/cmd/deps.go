package main

import (
	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/domain/service"
	pb "github.com/artem-benda/gophkeeper/client/internal/infrastructure/grpc"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/repository"
)

func mustCreateAppDependencies(c pb.GophKeeperServiceClient, passKey string) *application.AppDependencies {
	userRepository := repository.NewUserRepository(c)
	userService := service.NewUserService(userRepository)

	secretRepository := repository.NewSecretRepository(c)
	secretService := service.NewSecretService(secretRepository, passKey)
	return &application.AppDependencies{
		US: userService,
		SS: secretService,
	}
}
