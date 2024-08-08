package main

import (
	"github.com/artem-benda/gophkeeper/client/internal/application"
	"github.com/artem-benda/gophkeeper/client/internal/domain/service"
	pb "github.com/artem-benda/gophkeeper/client/internal/infrastructure/grpc"
	"github.com/artem-benda/gophkeeper/client/internal/infrastructure/repository"
)

func mustCreateAppDependencies(c pb.GophKeeperServiceClient) *application.AppDependencies {
	userRepository := repository.NewUserRepository(c)
	userService := service.NewUserService(userRepository)
	return &application.AppDependencies{
		US: userService,
	}
}
