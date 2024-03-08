package reports

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	rppb "parking/gunk/api/v1/reports"
)

type ReportCore interface {
	GetReport(ctx context.Context, req *rppb.GetReportRequest) (*rppb.GetReportResponse, error)
}

type Svc struct {
	rppb.UnimplementedReportServiceServer
	core ReportCore
}

func New(core ReportCore) *Svc {
	return &Svc{
		core: core,
	}
}

// RegisterSvc RegisterService with grpc server.
func (s *Svc) RegisterSvc(srv *grpc.Server) error {
	rppb.RegisterReportServiceServer(srv, s)
	return nil
}

// RegisterGateway grpcgw
func (s *Svc) RegisterGateway(ctx context.Context, mux *runtime.ServeMux, address string, options []grpc.DialOption) error {
	return rppb.RegisterReportServiceHandlerFromEndpoint(ctx, mux, address, options)
}
