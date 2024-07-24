package service

import (
	pb "auth-service/genproto/admin"
	"auth-service/models"
	"auth-service/pkg/logger"
	"auth-service/storage/postgres"
	"context"
	"database/sql"
	"log/slog"

	"github.com/pkg/errors"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	Repo    *postgres.AdminRepo
	RepoLoc *postgres.LocationRepo
	Logger  *slog.Logger
}

func NewAdminService(db *sql.DB) *AdminService {
	return &AdminService{
		Repo:    postgres.NewAdminRepo(db),
		RepoLoc: postgres.NewLocationRepo(db),
		Logger:  logger.NewLogger(),
	}
}

func (s *AdminService) AddUser(ctx context.Context, req *pb.NewUser) (*pb.NewUserResp, error) {
	s.Logger.Info("AddUser is starting")

	resp, err := s.Repo.Add(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to add user")
		s.Logger.Error(er.Error())
		return nil, er
	}

	_, err = s.RepoLoc.Add(ctx, &models.NewLocation{
		UserId:     resp.Id,
		Address:    req.Address,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	})
	if err != nil {
		er := errors.Wrap(err, "failed to add user location")
		s.Logger.Error(er.Error())
		return nil, er
	}

	s.Logger.Info("AddUser is finished")
	return resp, nil
}

func (s *AdminService) GetUser(ctx context.Context, req *pb.ID) (*pb.UserInfo, error) {
	s.Logger.Info("GetUser is starting")

	resp, err := s.Repo.Read(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to get user")
		s.Logger.Error(er.Error())
		return nil, er
	}

	loc, err := s.RepoLoc.Read(ctx, resp.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get user location")
		s.Logger.Error(er.Error())
		return nil, er
	}
	resp.Address = loc.Address
	resp.City = loc.City
	resp.State = loc.State
	resp.Country = loc.Country
	resp.PostalCode = loc.PostalCode

	s.Logger.Info("GetUser is finished")
	return resp, nil
}

func (s *AdminService) UpdateUser(ctx context.Context, req *pb.NewData) (*pb.NewDataResp, error) {
	s.Logger.Info("UpdateUser is starting")

	resp, err := s.Repo.Update(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to update user")
		s.Logger.Error(er.Error())
		return nil, er
	}

	_, err = s.RepoLoc.Update(ctx, &models.NewLocation{
		UserId:     resp.Id,
		Address:    req.Address,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	})
	if err != nil {
		er := errors.Wrap(err, "failed to update user location")
		s.Logger.Error(er.Error())
		return nil, er
	}

	s.Logger.Info("UpdateUser is finished")
	return resp, nil
}

func (s *AdminService) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	s.Logger.Info("DeleteUser is starting")

	err := s.Repo.Delete(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to delete user")
		s.Logger.Error(er.Error())
		return nil, er
	}

	err = s.RepoLoc.Delete(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to delete user location")
		s.Logger.Error(er.Error())
		return nil, er
	}

	s.Logger.Info("DeleteUser is finished")
	return &pb.Void{}, nil
}

func (s *AdminService) FetchUsers(ctx context.Context, req *pb.Filter) (*pb.Users, error) {
	s.Logger.Info("FetchUsers is starting")

	resp, err := s.Repo.FetchUsers(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to fetch users")
		s.Logger.Error(er.Error())
		return nil, er
	}

	s.Logger.Info("FetchUsers is finished")
	return resp, nil
}
