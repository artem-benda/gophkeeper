package server

import (
	"fmt"
	"github.com/artem-benda/gophkeeper/server/internal/application/grpc"
	pb "github.com/artem-benda/gophkeeper/server/internal/application/grpc"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	g "google.golang.org/grpc"
	"net"
)

func mustRunGrpcServer(dbpool *pgxpool.Pool, endpoint string) {
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		panic(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := g.NewServer()
	// регистрируем сервис
	pb.RegisterGophKeeperServiceServer(s, &grpc.MetricsGrpsServer{Storage: storage, DBPool: dbpool})

	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := s.Serve(listen); err != nil {

		logger.Log.Debug("server sut down", zap.Error(err))
		// Сбрасываем на диск данные из хранилища, только для memStorage
		if flushStorage != nil {
			err = flushStorage()
			if err != nil {
				logger.Log.Error("error flushing storage on shutdown")
			}
		}
	}
}
