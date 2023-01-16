package main

import (
	trippb "coolcar/proto/gen/go"
	trip "coolcar/tripservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	trippb.RegisterTripServiceServer(s, &trip.ServiceServer{})
	log.Fatal(s.Serve(lis))
}
