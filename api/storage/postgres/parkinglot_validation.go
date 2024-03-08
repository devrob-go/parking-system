package postgres

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"

	"parking/api/storage"
	"parking/utility/logging"
)

func (s *Storage) CreateParkingLotValidation(ctx context.Context, req *storage.CreateParkingLotRequest) error {
	log := logging.FromContext(ctx).WithField("method", "storage.CreateParkingLotValidation")

	if req == nil {
		return errors.New("request can not be empty")
	}

	required := validation.Required
	if err := validation.ValidateStruct(req,
		validation.Field(&req.Name, required, validation.Length(2, 50)),
		validation.Field(&req.TotalSpaces, required),
	); err != nil {
		logging.WithError(err, log).Error("failed to validate Create ParkingLot")
		return err
	}

	return nil
}
