package main

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/realtemirov/book_service/config"
	"github.com/realtemirov/book_service/grpc"
	"github.com/realtemirov/book_service/pkg/logger"
	"github.com/realtemirov/book_service/storage/postgres"
	"github.com/spf13/cast"
)

func main() {

	cfg := config.Load()

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {

	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelInfo
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	// connecting to db
	pgdb, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	// setting grpc server
	grpcServer := grpc.SetUpServer(cfg, log, pgdb)

	// listening port for req
	grpclistener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	log.Info("gRPC: Server is being started...", logger.String("port: ", cast.ToString(cfg.GRPCPort)))

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", cfg.HTTPPort))
	// serving to grpc
	if err := grpcServer.Serve(grpclistener); err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

}
