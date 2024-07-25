package service

import (
	pb "auth-service/genproto/user"
	"auth-service/models"
	"auth-service/pkg/logger"
	"auth-service/storage"
	"context"
	"log/slog"

	"github.com/pkg/errors"
)

type UserService struct {
	pb.UnimplementedUserServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewUserService(s storage.IStorage) *UserService {
	return &UserService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (r *UserService) GetProfile(ctx context.Context, req *pb.Void) (*pb.Profile, error) {
	r.logger.Info("GetUserProfile is starting")
	res, err := r.storage.User().GetProfile(ctx)
	if err != nil {
		er := errors.Wrap(err, "failed to get profile")
		r.logger.Error(er.Error())
		return nil, er
	}

	loc, err := r.storage.Location().Read(ctx, res.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get user location")
		r.logger.Error(er.Error())
		return nil, er
	}
	res.Address = loc.Address
	res.City = loc.City
	res.State = loc.State
	res.Country = loc.Country
	res.PostalCode = loc.PostalCode

	r.logger.Info("GetUserProfile has finished")
	return res, nil
}

func (r *UserService) UpdateProfile(ctx context.Context, req *pb.NewData) (*pb.UpdateResp, error) {
	r.logger.Info("UpdateUserProfile is starting")
	res, err := r.storage.User().UpdateProfile(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to update profile")
		r.logger.Error(er.Error())
		return nil, er
	}
	_, err = r.storage.Location().Update(ctx, &models.NewLocation{
		UserId:     res.Id,
		Address:    req.Address,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	})
	if err != nil {
		er := errors.Wrap(err, "failed to update user location")
		r.logger.Error(er.Error())
		return nil, er
	}

	r.logger.Info("UpdateUserProfile has finished")
	return res, nil
}

func (r *UserService) DeleteProfile(ctx context.Context, req *pb.Void) (*pb.Void, error) {
	r.logger.Info("DeleteProfile is starting")
	err := r.storage.User().DeleteProfile(ctx)
	if err != nil {
		er := errors.Wrap(err, "failed to delete profile")
		r.logger.Error(er.Error())
		return nil, er
	}
	Id := ctx.Value("user_id").(string)

	err = r.storage.Location().Delete(ctx, Id)
	if err != nil {
		er := errors.Wrap(err, "failed to delete user location")
		r.logger.Error(er.Error())
		return nil, er
	}
	r.logger.Info("DeleteProfile has finished")
	return &pb.Void{}, nil
}
