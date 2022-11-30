package grpc

import (
	"github.com/realtemirov/book_service/config"
	"github.com/realtemirov/book_service/genproto/book_service"
	"github.com/realtemirov/book_service/genproto/order_service"
	"github.com/realtemirov/book_service/grpc/service"
	"github.com/realtemirov/book_service/pkg/logger"
	"github.com/realtemirov/book_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	book_service.RegisterBookServiceServer(grpcServer, service.NewBookService(cfg, log, strg))
	order_service.RegisterOrderServiceServer(grpcServer, service.NewOrderService(cfg, log, strg))
	reflection.Register(grpcServer)

	return
}
