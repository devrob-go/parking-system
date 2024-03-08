package vehicle

import (
	"context"

	"google.golang.org/grpc/codes"
	rpc "google.golang.org/grpc/status"

	rppb "parking/gunk/api/v1/reports"
	"parking/utility/logging"
)

func (s *Svc) GetReport(ctx context.Context, req *rppb.GetReportRequest) (*rppb.GetReportResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "core.GetReport")
	log.Trace("request received")

	rData, err := s.store.GetReport(ctx, req.StartDate.AsTime(), req.GetEndDate().AsTime())
	if err != nil {
		logging.WithError(err, log).Error("failed to handle report response")
		return nil, rpc.Error(codes.Internal, "processing failed")
	}

	return &rppb.GetReportResponse{
		Data: &rppb.ReportData{
			Report: &rppb.Report{
				TotalEarning: rData.TotalEarning.Int32,
				TotalParked:  rData.TotalParked.Int32,
				TotalHours:   rData.TotalHours.Int32,
			},
		},
	}, nil
}
