package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"

	"parking/api/storage"
	"parking/utility/logging"
)

const insertParkVehicle = `
	INSERT INTO vehicle (
		license_plate, 
		entry_time,
		slot_id
	) VALUES (
		:license_plate, 
		:entry_time,
		:slot_id
	) RETURNING *
`

func (s *Storage) ParkVehicle(ctx context.Context, req *storage.ParkVehicleRequest) (*storage.VehicleResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "storage.ParkVehicle")
	if err := s.ParkVehicleValidation(ctx, req); err != nil {
		logging.WithError(err, log).Error("invalid request")
		return nil, storage.InvalidArgument
	}

	pstmt, err := s.db.PrepareNamedContext(ctx, insertParkVehicle)
	if err != nil {
		logging.WithError(err, log).Error("failed to prepare query")
		return nil, err
	}
	defer pstmt.Close()

	res := &storage.VehicleResponse{}
	if err := pstmt.Get(res, req); err != nil {
		pErr, ok := err.(*pq.Error)
		if ok && pErr.Code == pqUnique {
			logging.WithError(err, log).Error("failed to update parking slot to duplicate row")
			return nil, storage.Conflict
		}
		logging.WithError(err, log).Error("failed to update parking slot")
		return nil, err
	}

	return res, nil
}

const updateUnParkingSlot = `
	UPDATE vehicle SET
		exit_time 	 = :exit_time,
		parked_hours = :parked_hours,
		fee 		 = :fee,
		updated_at 	 = :updated_at
	WHERE id = :id
	RETURNING *
`

func (s *Storage) UnParkVehicle(ctx context.Context, req *storage.UnParkVehicleRequest) (*storage.VehicleResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "storage.UnParkVehicle")
	if err := s.UnParkVehicleValidation(ctx, req); err != nil {
		logging.WithError(err, log).Error("invalid request")
		return nil, storage.InvalidArgument
	}

	pstmt, err := s.db.PrepareNamedContext(ctx, updateUnParkingSlot)
	if err != nil {
		logging.WithError(err, log).Error("failed to prepare query")
		return nil, err
	}
	defer pstmt.Close()

	res := &storage.VehicleResponse{}
	if err := pstmt.Get(res, req); err != nil {
		pErr, ok := err.(*pq.Error)
		if ok && pErr.Code == pqUnique {
			logging.WithError(err, log).Error("failed to update unparking slot to duplicate row")
			return nil, storage.Conflict
		}
		logging.WithError(err, log).Error("failed to update unparking slot")
		return nil, err
	}

	return res, nil
}

func (s *Storage) GetVehicleById(ctx context.Context, id string) (*storage.VehicleResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "storage.GetVehicleById")
	if id == "" {
		return nil, errors.New("vehicle id can not be blank")
	}

	getVehicle := `SELECT v.*, ps.lot_id, ps.slot_number FROM vehicle v LEFT JOIN parking_slot ps ON ps.id = v.slot_id WHERE v.id = $1`
	var vehicleData storage.VehicleResponse
	if err := s.db.GetContext(ctx, &vehicleData, getVehicle, id); err != nil {
		logging.WithError(err, log).Error("failed to get parking details in storage")
		if err == sql.ErrNoRows {
			return nil, storage.NotFound
		}
		return nil, err
	}

	return &vehicleData, nil
}
