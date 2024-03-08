package parkingslot

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"parking/api/storage"
	pspb "parking/gunk/api/v1/parkingslot"
)

// To handle single object response
func handleParkingSlotResponse(res *storage.ParkingSlotResponse) (*pspb.ParkingSlot, error) {
	if res.ID == "" {
		return nil, status.Error(codes.NotFound, "no data found")
	}
	return handler(append([]*storage.ParkingSlotResponse{}, res))[0], nil
}

// To handle list response
func handleListParkingSlotResponse(res []*storage.ParkingSlotResponse) ([]*pspb.ParkingSlot, error) {
	if res == nil {
		return nil, status.Error(codes.NotFound, "no data found")
	}
	return handler(res), nil
}

// Common handler for single/list response of bank accounts
func handler(res []*storage.ParkingSlotResponse) []*pspb.ParkingSlot {
	ParkingSlot := make([]*pspb.ParkingSlot, 0, len(res))

	for _, v := range res {
		ParkingSlot = append(ParkingSlot, &pspb.ParkingSlot{
			ID:         v.ID,
			LotID:      v.LotID,
			SlotNumber: v.SlotNumber,
			Status:     pspb.Status(pspb.Status_value[v.Status.String()]),
			CreatedAt:  timestamppb.New(v.CreatedAt),
			UpdatedAt:  timestamppb.New(v.UpdatedAt),
		})
	}

	if len(ParkingSlot) == 0 {
		return nil
	}

	return ParkingSlot
}
