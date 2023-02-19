package server

import (
	"coolcar/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

// GRPCConfig a grpc server
type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(c *GRPCConfig) error {
	nameFiled := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameFiled, zap.Error(err))
	}
	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			c.Logger.Fatal("cannot create auth interceptor", nameFiled, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)

	return s.Serve(lis)
}
