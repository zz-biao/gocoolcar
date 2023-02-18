package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	jsonpb := &runtime.JSONPb{}
	jsonpb.UseEnumNumbers = true
	jsonpb.UseProtoNames = true
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb))
	//auth 服务
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(c, mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("connot register auth service: %v", err)
	}
	//trip 服务
	err = rentalpb.RegisterTripServiceHandlerFromEndpoint(c, mux, "localhost:8082", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("connot register auth service: %v", err)
	}
	log.Fatal(http.ListenAndServe(":8080", mux))
}
