package parkinglot

import (
	"context"
	"math"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"parking/api/storage"
	plpb "parking/gunk/api/v1/parkinglot"
	"parking/utility/logging"
)

func (s *Svc) CreateParkingLot(ctx context.Context, req *plpb.CreateParkingLotRequest) (*plpb.CreateParkingLotResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "core.CreateParkingLot")
	log.Trace("request received")

	res, err := s.store.CreateParkingLot(ctx, &storage.CreateParkingLotRequest{
		Name:        req.GetName(),
		TotalSpaces: req.GetTotalSpace(),
	})
	if err != nil {
		logging.WithError(err, log).Error("failed to create ParkingLot")
		if err == storage.NotFound {
			return nil, status.Error(codes.NotFound, "no data found")
		}
		return nil, status.Error(codes.Internal, "processing failed")
	}

	aRes, err := handleParkingLotResponse(res)
	if err != nil {
		logging.WithError(err, log).Error("failed to handle ParkingLot response")
		return nil, status.Error(codes.Internal, "processing failed")
	}

	return &plpb.CreateParkingLotResponse{
		Data: &plpb.ParkingLotData{
			ParkingLot: aRes,
		},
	}, nil
}

func (s *Svc) ListParkingLot(ctx context.Context, req *plpb.ListParkingLotRequest) (*plpb.ListParkingLotResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "core.ListParkingLot")
	log.Trace("request received")

	res, err := s.store.ListParkingLot(ctx, &storage.ListParkingLotRequest{
		OrderBy:      storage.OrderBy(req.OrderBy.String()),
		SearchTerm:   req.GetSearchTerm(),
		SortByColumn: storage.SortByColumn(req.GetSortByColumn()),
		Limit:        int(req.GetLimit()),
		Offset:       int(req.GetOffset()),
	})
	if err != nil || res == nil {
		logging.WithError(err, log).Error("failed to get ParkingLot List")
		if err == storage.NotFound {
			return nil, status.Error(codes.NotFound, "no data found")
		}
		return nil, status.Error(codes.Internal, "processing failed")
	}
	if len(res) == 0 {
		return &plpb.ListParkingLotResponse{
			Data: &plpb.ListParkingLotData{
				ParkingLot: nil,
			},
			Total: 0,
		}, nil
	}
	aRes, err := handleListParkingLotResponse(res)
	if err != nil {
		logging.WithError(err, log).Error("failed to handle list ParkingLot response")
		return nil, status.Error(codes.Internal, "processing failed")
	}

	var totalParkingLot int32 = 0
	if len(res) > 0 {
		totalParkingLot = res[0].Total
	}

	return &plpb.ListParkingLotResponse{
		Data: &plpb.ListParkingLotData{
			ParkingLot: aRes,
		},
		Total: totalParkingLot,
	}, nil
}

func (s *Svc) GetParkingLot(ctx context.Context, req *plpb.GetParkingLotRequest) (*plpb.GetParkingLotResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "core.GetParkingLot")
	log.Trace("request received")

	res, err := s.store.GetParkinglotDetailsByID(ctx, req.GetID())
	if err != nil || res == nil {
		logging.WithError(err, log).Error("failed to get parking lot")
		if err == storage.NotFound {
			return nil, status.Error(codes.NotFound, "no data found")
		}
		return nil, status.Error(codes.Internal, "processing failed")
	}
	parkedVehicles := make(map[string][]*plpb.ParkedVehicles, 0)
	for _, v := range res {
		parkedHours := int32(math.Ceil(time.Now().Sub(v.EntryTime).Hours()))
		parkedVehicles[v.ID] = append(parkedVehicles[v.ID], &plpb.ParkedVehicles{
			LicensePlate: v.LicensePlate,
			EntryTime:    timestamppb.New(v.EntryTime),
			ParkedHours:  parkedHours,
			Fee:          parkedHours * FEE_PER_HOUR,
			SlotID:       v.SlotID,
		})
	}
	var parkingLot *plpb.GetParkingLot
	if len(res) > 0 {
		parkingLot = &plpb.GetParkingLot{
			ID:             res[0].ID,
			Name:           res[0].Name,
			TotalSpace:     res[0].TotalSpace,
			TotalParked:    res[0].TotalParked,
			ParkedVehicles: parkedVehicles[res[0].ID],
			CreatedAt:      timestamppb.New(res[0].CreatedAt),
			UpdatedAt:      timestamppb.New(res[0].UpdatedAt),
		}
	}
	return &plpb.GetParkingLotResponse{
		Data: &plpb.GetParkingLotData{
			ParkingLot: parkingLot,
		},
	}, nil
}
