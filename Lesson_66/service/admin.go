package service

import (
	pb "auth-service/genproto/admin"
	"auth-service/models"
	"auth-service/pkg/logger"
	"auth-service/storage"
	"context"
	"log/slog"

	"github.com/pkg/errors"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewAdminService(s storage.IStorage) *AdminService {
	return &AdminService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (s *AdminService) AddUser(ctx context.Context, req *pb.NewUser) (*pb.NewUserResp, error) {
	s.logger.Info("AddUser is starting")

	resp, err := s.storage.Admin().Add(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to add user")
		s.logger.Error(er.Error())
		return nil, er
	}

	_, err = s.storage.Location().Add(ctx, &models.NewLocation{
		UserId:     resp.Id,
		Address:    req.Address,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	})
	if err != nil {
		er := errors.Wrap(err, "failed to add user location")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("AddUser is finished")
	return resp, nil
}

func (s *AdminService) GetUser(ctx context.Context, req *pb.ID) (*pb.UserInfo, error) {
	s.logger.Info("GetUser is starting")

	resp, err := s.storage.Admin().Read(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to get user")
		s.logger.Error(er.Error())
		return nil, er
	}

	loc, err := s.storage.Location().Read(ctx, resp.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get user location")
		s.logger.Error(er.Error())
		return nil, er
	}
	resp.Address = loc.Address
	resp.City = loc.City
	resp.State = loc.State
	resp.Country = loc.Country
	resp.PostalCode = loc.PostalCode

	s.logger.Info("GetUser is finished")
	return resp, nil
}

func (s *AdminService) UpdateUser(ctx context.Context, req *pb.NewData) (*pb.NewDataResp, error) {
	s.logger.Info("UpdateUser is starting")

	resp, err := s.storage.Admin().Update(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to update user")
		s.logger.Error(er.Error())
		return nil, er
	}

	_, err = s.storage.Location().Update(ctx, &models.NewLocation{
		UserId:     resp.Id,
		Address:    req.Address,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	})
	if err != nil {
		er := errors.Wrap(err, "failed to update user location")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("UpdateUser is finished")
	return resp, nil
}

func (s *AdminService) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	s.logger.Info("DeleteUser is starting")

	err := s.storage.Admin().Delete(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to delete user")
		s.logger.Error(er.Error())
		return nil, er
	}

	err = s.storage.Location().Delete(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to delete user location")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("DeleteUser is finished")
	return &pb.Void{}, nil
}

func (s *AdminService) FetchUsers(ctx context.Context, req *pb.Filter) (*pb.Users, error) {
	s.logger.Info("FetchUsers is starting")

	resp, err := s.storage.Admin().FetchUsers(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to fetch users")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("FetchUsers is finished")
	return resp, nil
}
