package main

import (
	"fmt"
	"log/slog"
	"net"

	pb "github.com/artem-benda/gophkeeper/server/internal/application/grpc"
	"github.com/artem-benda/gophkeeper/server/internal/application/middleware"
	"github.com/artem-benda/gophkeeper/server/internal/application/server"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	g "google.golang.org/grpc"
)

// mustRunGrpcServer - запустить GRPC сервер
func mustRunGrpcServer(userService contract.UserService, secretService contract.SecretService, cfg Config) {
	listen, err := net.Listen("tcp", cfg.Endpoint)
	if err != nil {
		panic(err)
	}
	// создаём gRPC-сервер с перехватчиком для авторизации пользователя
	s := g.NewServer(
		g.UnaryInterceptor(auth.UnaryServerInterceptor(middleware.DummyAuthFunc)),
	)
	// регистрируем сервис
	pb.RegisterGophKeeperServiceServer(s, &server.GophKeeperGrpcServer{USvc: userService, SSvc: secretService})

	fmt.Println("Сервер gRPC начал работу...")
	if err := s.Serve(listen); err != nil {
		slog.Error("server sut down", slog.Any("error", err))
	}
}
