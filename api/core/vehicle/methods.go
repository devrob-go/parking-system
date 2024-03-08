package vehicle

import (
	"context"
	"math"
	"time"

	"google.golang.org/grpc/codes"
	rpc "google.golang.org/grpc/status"

	"parking/api/storage"
	pvpb "parking/gunk/api/v1/vehicle"
	"parking/utility/logging"
)

func (s *Svc) ParkVehicle(ctx context.Context, req *pvpb.ParkVehicleRequest) (*pvpb.ParkVehicleResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "core.ParkVehicle")
	log.Trace("request received")

	slots, err := s.store.GetParkingSlotsByLotID(ctx, req.GetLotID())
	if err != nil {
		logging.WithError(err, log).Error("failed to handle parking lot response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	slotNumber := int32(1)
	if len(slots) > 0 { // to find out the lowest available slot
		totalSpace := int(slots[0].TotalSpace)
		numbers := make([]int, totalSpace)
		for i := 1; i <= totalSpace; i++ {
			numbers[i-1] = i
		}

		// Numbers to drop
		toDrop := []int{}
		occupied := storage.SlotStatusOccupied.String()
		maintenance := storage.SlotStatusMaintenance.String()
		for _, v := range slots {
			if v.Status.String() == occupied || v.Status.String() == maintenance {
				toDrop = append(toDrop, int(v.SlotNumber))
			}
		}

		// Remove specified numbers from the slice
		for _, drop := range toDrop {
			for i, num := range numbers {
				if num == drop {
					numbers = append(numbers[:i], numbers[i+1:]...)
					break
				}
			}
		}

		// Find the minimum value in the remaining slice
		minimum := math.MaxInt
		for _, num := range numbers {
			if num < minimum {
				minimum = num
			}
		}
		slotNumber = int32(minimum)
	}

	slot, err := s.store.UpdateParkingSlot(ctx, &storage.UpdateParkingSlotRequest{
		LotID:      req.GetLotID(),
		SlotNumber: slotNumber,
		Status:     storage.SlotStatusOccupied,
	})
	if err != nil {
		logging.WithError(err, log).Error("failed to handle parking slot response")
		return nil, rpc.Error(codes.Internal, "failed to park the vehicle")
	}

	res, err := s.store.ParkVehicle(ctx, &storage.ParkVehicleRequest{
		LicensePlate: req.GetLicensePlate(),
		EntryTime:    time.Now(),
		SlotID:       slot.ID,
	})
	if err != nil {
		logging.WithError(err, log).Error("failed to update parking slot")
		if err == storage.NotFound {
			return nil, rpc.Error(codes.NotFound, "no data found")
		}
		return nil, rpc.Error(codes.Internal, "failed to park the vehicle")
	}

	vRes, err := handleVehicleResponse(res)
	if err != nil {
		logging.WithError(err, log).Error("failed to handle parking slot response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	return &pvpb.ParkVehicleResponse{
		Data: &pvpb.VehicleData{
			Vehicle: vRes,
		},
	}, nil
}

func (s *Svc) UnParkVehicle(ctx context.Context, req *pvpb.UnParkVehicleRequest) (*pvpb.UnParkVehicleResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "core.UnParkVehicle")
	log.Trace("request received")

	vData, err := s.store.GetVehicleById(ctx, req.GetID())
	if err != nil {
		logging.WithError(err, log).Error("failed to handle vehicle response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	parkedHours := int32(math.Ceil(time.Now().Sub(vData.EntryTime).Hours()))

	upvData, err := s.store.UnParkVehicle(ctx, &storage.UnParkVehicleRequest{
		ID:          req.GetID(),
		ExitTime:    time.Now(),
		ParkedHours: parkedHours,
		Fee:         parkedHours * FEE_PER_HOUR,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		logging.WithError(err, log).Error("failed to handle vehicle response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	if _, err := s.store.UpdateParkingSlot(ctx, &storage.UpdateParkingSlotRequest{
		LotID:      vData.LotID,
		SlotNumber: vData.SlotNumber,
		Status:     storage.SlotStatusAvailable,
	}); err != nil {
		logging.WithError(err, log).Error("failed to handle parking slot response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	vRes, err := handleVehicleResponse(upvData)
	if err != nil {
		logging.WithError(err, log).Error("failed to handle parking slot response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	return &pvpb.UnParkVehicleResponse{
		Data: &pvpb.VehicleData{
			Vehicle: vRes,
		},
	}, nil
}
