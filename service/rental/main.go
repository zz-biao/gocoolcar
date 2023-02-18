package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	s := grpc.NewServer()
	rentalpb.RegisterTripServiceServer(s, &trip.Service{ //设置配置
		Logger: logger,
	})
	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("cannot server", zap.Error(err))
	}
}
