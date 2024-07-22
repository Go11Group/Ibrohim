package service

import (
	pb "auth-service/genproto/user"
	"auth-service/pkg/logger"
	"auth-service/storage/postgres"
	"context"
	"database/sql"
	"log/slog"

	"github.com/pkg/errors"
)

type UserService struct {
	pb.UnimplementedUserServer
	Repo   *postgres.UserRepo
	Logger *slog.Logger
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		Repo:   postgres.NewUserRepo(db),
		Logger: logger.NewLogger(),
	}
}

func (r *UserService) GetProfile(ctx context.Context, req *pb.Void) (*pb.Profile, error) {
	r.Logger.Info("GetUserProfile is starting")
	res, err := r.Repo.GetProfile(ctx)
	if err != nil {
		er := errors.Wrap(err, "failed to get profile")
		r.Logger.Error(er.Error())
		return nil, er
	}
	r.Logger.Info("GetUserProfile has finished")
	return res, nil
}

func (r *UserService) UpdateProfile(ctx context.Context, req *pb.NewData) (*pb.UpdateResp, error) {
	r.Logger.Info("UpdateUserProfile is starting")
	res, err := r.Repo.UpdateProfile(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to update profile")
		r.Logger.Error(er.Error())
		return nil, er
	}
	r.Logger.Info("UpdateUserProfile has finished")
	return res, nil
}

func (r *UserService) DeleteProfile(ctx context.Context, req *pb.Void) (*pb.Void, error) {
	r.Logger.Info("DeleteProfile is starting")
	err := r.Repo.DeleteProfile(ctx)
	if err != nil {
		er := errors.Wrap(err, "failed to delete profile")
		r.Logger.Error(er.Error())
		return nil, er
	}
	r.Logger.Info("DeleteProfile has finished")
	return &pb.Void{}, nil
}
