package parkinglot

import (
	"context"

	"github.com/sirupsen/logrus"

	"parking/api/storage"
)

const FEE_PER_HOUR = int32(10)

type ParkingLotStore interface {
	CreateParkingLot(ctx context.Context, req *storage.CreateParkingLotRequest) (*storage.ParkingLotResponse, error)
	ListParkingLot(ctx context.Context, req *storage.ListParkingLotRequest) ([]*storage.ParkingLotResponse, error)
	GetParkinglotDetailsByID(ctx context.Context, lotID string) ([]*storage.ParkingLotDetailsResponse, error)
}

type Svc struct {
	store  ParkingLotStore
	logger *logrus.Entry
}

func New(store ParkingLotStore, logger *logrus.Entry) *Svc {
	return &Svc{
		store:  store,
		logger: logger,
	}
}
