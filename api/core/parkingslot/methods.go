package parkingslot

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	rpc "google.golang.org/grpc/status"

	"parking/api/storage"
	pspb "parking/gunk/api/v1/parkingslot"
	"parking/utility/logging"
)

func (s *Svc) UpdateParkingSlot(ctx context.Context, req *pspb.UpdateParkingSlotRequest) (*pspb.UpdateParkingSlotResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "core.UpdateParkingSlot")
	log.Trace("request received")

	if req.GetStatus().String() == pspb.Status_UnknownStatus.String() {
		return nil, rpc.Error(codes.InvalidArgument, "invalid status")
	}

	lot, err := s.store.GetParkinglotByID(ctx, req.GetLotID())
	if err != nil {
		logging.WithError(err, log).Error("failed to handle parking lot response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	if req.GetSlotNumber() > lot.TotalSpace {
		logging.WithError(errors.New("failed to handle parking lot response"), log)
		return nil, rpc.Error(codes.InvalidArgument, "slot number can not be bigger than the space")
	}

	res, err := s.store.UpdateParkingSlot(ctx, &storage.UpdateParkingSlotRequest{
		LotID:      req.GetLotID(),
		SlotNumber: req.GetSlotNumber(),
		Status:     storage.SlotStatus(req.GetStatus().String()),
	})
	if err != nil {
		logging.WithError(err, log).Error("failed to update parking slot")
		if err == storage.NotFound {
			return nil, rpc.Error(codes.NotFound, "no data found")
		}
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	lRes, err := handleParkingSlotResponse(res)
	if err != nil {
		logging.WithError(err, log).Error("failed to handle parking slot response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	return &pspb.UpdateParkingSlotResponse{
		Data: &pspb.ParkingSlotData{
			ParkingSlot: lRes,
		},
	}, nil
}
