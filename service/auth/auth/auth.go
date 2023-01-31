package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	Logger         *zap.Logger
}

type OpenIDResolver interface {
	Resolver(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (res *authpb.LoginResponse, err error) {

	st, err := s.OpenIDResolver.Resolver(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}
	log.Printf(st)

	s.Logger.Info("received code", zap.String("code", req.Code))
	return &authpb.LoginResponse{
		AccessToken: "123123",
		ExpiresIn:   3600,
	}, nil
}
