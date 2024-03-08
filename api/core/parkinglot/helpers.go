package parkinglot

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"parking/api/storage"
	plpb "parking/gunk/api/v1/parkinglot"
)

// To handle single object response
func handleParkingLotResponse(res *storage.ParkingLotResponse) (*plpb.ParkingLot, error) {
	if res.ID == "" {
		return nil, status.Error(codes.NotFound, "no data found")
	}
	return handler(append([]*storage.ParkingLotResponse{}, res))[0], nil
}

// To handle list response
func handleListParkingLotResponse(res []*storage.ParkingLotResponse) ([]*plpb.ParkingLot, error) {
	if res == nil {
		return nil, status.Error(codes.NotFound, "no data found")
	}
	return handler(res), nil
}

// Common handler for single/list response of bank accounts
func handler(res []*storage.ParkingLotResponse) []*plpb.ParkingLot {
	parkinglot := make([]*plpb.ParkingLot, 0, len(res))

	for _, v := range res {
		parkinglot = append(parkinglot, &plpb.ParkingLot{
			ID:         v.ID,
			Name:       v.Name,
			TotalSpace: v.TotalSpace,
			CreatedAt:  timestamppb.New(v.CreatedAt),
			UpdatedAt:  timestamppb.New(v.UpdatedAt),
		})
	}

	if len(parkinglot) == 0 {
		return nil
	}

	return parkinglot
}
