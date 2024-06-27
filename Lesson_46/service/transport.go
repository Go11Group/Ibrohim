package service

import (
	"context"
	"database/sql"
	pb "weather-transport/genproto/transport"
	"weather-transport/storage/postgres"

	"github.com/pkg/errors"
)

type transportService struct {
	pb.UnimplementedTransportServiceServer
	Repo *postgres.BusRepo
}

func NewTransportService(db *sql.DB) *transportService {
	return &transportService{Repo: postgres.NewBusRepo(db)}
}

func (t *transportService) GetBusSchedule(ctx context.Context, n *pb.Number) (*pb.Schedule, error) {
	if n.Number == "" {
		return nil, errors.New("empty field provided")
	}

	resp, err := t.Repo.GetBusSchedule(n)
	if err != nil {
		return nil, errors.Wrap(err, "error getting bus schedule")
	}

	return resp, nil
}

func (t *transportService) TrackBusLocation(ctx context.Context, n *pb.Number) (*pb.GPS, error) {
	if n.Number == "" {
		return nil, errors.New("empty field provided")
	}

	resp, err := t.Repo.TrackBusLocation(n)
	if err != nil {
		return nil, errors.Wrap(err, "error tracking bus location")
	}

	return resp, nil
}

func (t *transportService) ReportTrafficJam(ctx context.Context, r *pb.Route) (*pb.Traffic, error) {
	if r.Name == "" || r.Transports < 1 {
		return nil, errors.New("empty or invalid fields provided")
	}

	resp, err := t.Repo.ReportTrafficJam(r)
	if err != nil {
		return nil, errors.Wrap(err, "error reporting traffic jam")
	}

	return resp, nil
}
