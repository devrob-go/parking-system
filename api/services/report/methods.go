package reports

import (
	"context"

	rppb "parking/gunk/api/v1/reports"
	uErr "parking/utility/error/error"
	"parking/utility/logging"
)

func (s *Svc) GetReport(ctx context.Context, req *rppb.GetReportRequest) (*rppb.GetReportResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "service.GetReport")
	log.Trace("request received")
	res, err := s.core.GetReport(ctx, req)
	if err != nil {
		errM := "failed to park vehicle in core"
		log.WithError(err).Error(errM)
		return nil, uErr.HandleServiceErr(err)
	}

	return res, nil
}
