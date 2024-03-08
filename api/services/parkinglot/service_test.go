package parkinglot

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/testing/protocmp"

	plCore "parking/api/core/parkinglot"
	"parking/api/storage/postgres"
	plpb "parking/gunk/api/v1/parkinglot"
	"parking/utility/logging"
)

var _testStorage *postgres.Storage

func newTestSvc(t *testing.T, st *postgres.Storage) *Svc {
	conn := os.Getenv("DATABASE_CONNECTION")
	if conn == "" {
		t.Skip("missing database connection")
	}

	st, cleanup := postgres.NewTestStorage(conn, filepath.Join("..", "..", "migrations"))
	t.Cleanup(cleanup)

	logger := logging.NewLogger(nil).WithFields(logrus.Fields{
		"service": "test_parkinglot",
	})

	return New(plCore.New(st, logger))
}

func newTestStorage(tb testing.TB) *postgres.Storage {
	if testing.Short() {
		tb.Skip("skipping tests that use postgres on -short")
	}
	return _testStorage
}

type ParkingLotTestStruct struct {
	methodName string
	desc       string
	in         interface{}
	want       interface{}
	tops       cmp.Options
	wantErr    bool
}

var id string

func TestParkingLot(t *testing.T) {
	st := newTestStorage(t)
	s := newTestSvc(t, st)

	// NOTE: Can be added scenarios here to test more acceptance criteria
	tests := []ParkingLotTestStruct{
		{
			methodName: "CREATE",
			desc:       "Success Create ParkingLot",
			wantErr:    false,
			tops: cmp.Options{
				protocmp.Transform(),
			},
			in: &plpb.CreateParkingLotRequest{
				Name:       "Test Parking Lot",
				TotalSpace: 50,
			},
			want: &plpb.CreateParkingLotResponse{
				Data: &plpb.ParkingLotData{
					ParkingLot: &plpb.ParkingLot{
						Name:       "Test Parking Lot",
						TotalSpace: 50,
					},
				},
			},
		},
		{
			methodName: "GET_ALL",
			desc:       "Success Get ALL ParkingLot",
			wantErr:    false,
			tops: cmp.Options{
				cmpopts.IgnoreFields(plpb.ParkingLot{}, "ID", "CreatedAt", "UpdatedAt"),
				protocmp.Transform(),
			},
			in: &plpb.ListParkingLotRequest{},
			want: &plpb.ListParkingLotResponse{
				Data: &plpb.ListParkingLotData{
					ParkingLot: []*plpb.ParkingLot{
						{
							Name:       "Test Parking Lot",
							TotalSpace: 50,
						},
					},
				},
				Total: 1,
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.methodName, func(t *testing.T) {
			switch test.methodName {
			case "CREATE":
				CreateParkingLotTest(t, test, s)
			case "GET_ALL":
				ListParkingLotTest(t, test, s)
			}
		})
	}
}

func CreateParkingLotTest(t *testing.T, test ParkingLotTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*plpb.CreateParkingLotRequest)
	if !ok {
		t.Error("reqest type conversion error")
	}
	got, err := s.CreateParkingLot(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	id = got.Data.ParkingLot.ID
	o := test.tops
	if !cmp.Equal(test.want, got, o) {
		t.Error("(-want +got): ", cmp.Diff(test.want, got, o))
	}
}

func ListParkingLotTest(t *testing.T, test ParkingLotTestStruct, s *Svc) {
	ctx := context.Background()
	req, ok := test.in.(*plpb.ListParkingLotRequest)
	if !ok {
		t.Error("reqest type conversion error")
	}

	got, err := s.ListParkingLot(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	o := test.tops
	if !cmp.Equal(test.want, got, o) {
		t.Error("(-want +got): ", cmp.Diff(test.want, got, o))
	}
}
