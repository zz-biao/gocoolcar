package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

//type TripServiceServer interface {
//	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
//	mustEmbedUnimplementedTripServiceServer()
//}

// ServiceServer a trip service.
type ServiceServer struct{}

func (s *ServiceServer) mustEmbedUnimplementedTripServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (*ServiceServer) GetTrip(c context.Context, req *trippb.GetTripRequest) (res *trippb.GetTripResponse, err error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			Status:      trippb.TripStatus_FINISHED,
		},
	}, nil
}
