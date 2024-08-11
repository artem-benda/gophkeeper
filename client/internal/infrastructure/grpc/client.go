package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MustCreateGRPCClient создать клиент для gRPC сервера
func MustCreateGRPCClient(endpoint string) (GophKeeperServiceClient, *grpc.ClientConn) {
	// устанавливаем соединение с сервером
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// получаем переменную интерфейсного типа MonitorServiceClient,
	// через которую будем отправлять сообщения
	c := NewGophKeeperServiceClient(conn)
	return c, conn
}
