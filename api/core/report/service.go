package vehicle

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"parking/api/storage"
)

type ReportStore interface {
	GetReport(ctx context.Context, startDate, endDate time.Time) (*storage.Report, error)
}

type Svc struct {
	store  ReportStore
	logger *logrus.Entry
}

func New(store ReportStore, logger *logrus.Entry) *Svc {
	return &Svc{
		store:  store,
		logger: logger,
	}
}
