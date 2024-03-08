package vehicle

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"parking/api/storage"
	pvpb "parking/gunk/api/v1/vehicle"
)

// To handle single object response
func handleVehicleResponse(res *storage.VehicleResponse) (*pvpb.Vehicle, error) {
	if res.ID == "" {
		return nil, status.Error(codes.NotFound, "no data found")
	}
	return handler(append([]*storage.VehicleResponse{}, res))[0], nil
}

// To handle list response
func handleListVehicleResponse(res []*storage.VehicleResponse) ([]*pvpb.Vehicle, error) {
	if res == nil {
		return nil, status.Error(codes.NotFound, "no data found")
	}
	return handler(res), nil
}

// Common handler for single/list response of bank accounts
func handler(res []*storage.VehicleResponse) []*pvpb.Vehicle {
	Vehicle := make([]*pvpb.Vehicle, 0, len(res))
	for _, v := range res {
		exitTime := &timestamppb.Timestamp{}
		if v.ExitTime.Valid {
			exitTime = timestamppb.New(v.ExitTime.Time)
		}
		Vehicle = append(Vehicle, &pvpb.Vehicle{
			ID:           v.ID,
			LicensePlate: v.LicensePlate,
			EntryTime:    timestamppb.New(v.EntryTime),
			ExitTime:     exitTime,
			ParkedHours:  v.ParkedHours,
			Fee:          v.Fee,
			SlotID:       v.SlotID,
			CreatedAt:    timestamppb.New(v.CreatedAt),
			UpdatedAt:    timestamppb.New(v.UpdatedAt),
		})
	}

	if len(Vehicle) == 0 {
		return nil
	}

	return Vehicle
}
