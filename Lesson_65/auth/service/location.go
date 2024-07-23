package service

import (
	pb "auth-service/genproto/location"
	"auth-service/pkg/logger"
	"auth-service/storage/postgres"
	"context"
	"database/sql"
	"log/slog"

	"github.com/pkg/errors"
)

type LocationService struct {
	pb.UnimplementedLocationServer
	Repo *postgres.LocationRepo
	Log  *slog.Logger
}

func NewLocationService(db *sql.DB) *LocationService {
	return &LocationService{
		Repo: postgres.NewLocationRepo(db),
		Log:  logger.NewLogger(),
	}
}

func (s *LocationService) Add(ctx context.Context, req *pb.NewLocation) (*pb.NewLocationResp, error) {
	s.Log.Info("AddLocation is starting")

	resp, err := s.Repo.Add(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to add location")
		s.Log.Error(er.Error())
		return nil, er
	}

	s.Log.Info("AddLocation is finished")
	return resp, nil
}

func (s *LocationService) Get(ctx context.Context, req *pb.ID) (*pb.LocationDetails, error) {
	s.Log.Info("GetLocation is starting")

	resp, err := s.Repo.Read(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to get location")
		s.Log.Error(er.Error())
		return nil, er
	}

	s.Log.Info("GetLocation is finished")
	return resp, nil
}
