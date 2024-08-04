package main

import (
	"fmt"
	"log/slog"
	"net"

	pb "github.com/artem-benda/gophkeeper/server/internal/application/grpc"
	"github.com/artem-benda/gophkeeper/server/internal/application/server"
	"github.com/artem-benda/gophkeeper/server/internal/domain/contract"
	g "google.golang.org/grpc"
)

func mustRunGrpcServer(userService contract.UserService, secretService contract.SecretService, cfg Config) {
	listen, err := net.Listen("tcp", cfg.Endpoint)
	if err != nil {
		panic(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := g.NewServer()
	// регистрируем сервис
	pb.RegisterGophKeeperServiceServer(s, &server.GophKeeperGrpcServer{USvc: userService, SSvc: secretService})

	fmt.Println("Сервер gRPC начал работу...")
	if err := s.Serve(listen); err != nil {
		slog.Error("server sut down", slog.Any("error", err))
	}
}
