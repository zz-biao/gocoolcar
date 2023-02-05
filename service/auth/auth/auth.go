package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	Mongo          *dao.Mongo
	Logger         *zap.Logger
}

type OpenIDResolver interface {
	Resolver(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (res *authpb.LoginResponse, err error) {

	openID, err := s.OpenIDResolver.Resolver(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}

	accountID, err := s.Mongo.ResolveAccountID(c, openID)
	if err != nil {
		s.Logger.Error("cannot resolve account id ", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	s.Logger.Info("received code", zap.String("code", req.Code))
	return &authpb.LoginResponse{
		AccessToken: "token for openID:" + accountID,
		ExpiresIn:   3600,
	}, nil
}
