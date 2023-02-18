package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (rep *rentalpb.CreateTripResponse, err error) {
	s.Logger.Info("create trip", zap.String("start", req.Start))
	return nil, status.Error(codes.Unimplemented, "")
}
