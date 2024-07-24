package service

import (
	"auth-service/api/models"
	"auth-service/api/tokens"
	"auth-service/pkg/logger"
	"auth-service/storage/postgres"
	"context"
	"database/sql"
	"log/slog"

	"github.com/pkg/errors"
)

type AuthService struct {
	Repo   *postgres.UserRepo
	Logger *slog.Logger
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		Repo:   postgres.NewUserRepo(db),
		Logger: logger.NewLogger(),
	}
}

func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.Tokens, error) {
	s.Logger.Info("Login method is starting")

	id, username, passwordHash, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		er := errors.Wrap(err, "failed to find user")
		s.Logger.Error(er.Error())
		return nil, er
	}

	if passwordHash != req.Password {
		er := errors.New("incorrect password")
		s.Logger.Error(er.Error())
		return nil, er
	}

	accessToken, err := tokens.GenerateAccessToken(id, username, req.Email)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	refreshToken, err := tokens.GenerateRefreshToken(id)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	s.Logger.Info("Login has successfully finished")
	return &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, rt *models.RefreshToken) (*models.Tokens, error) {
	s.Logger.Info("CheckRefreshToken method is starting")

	_, err := tokens.ValidateRefreshToken(rt.Token)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	id, err := tokens.GetUserIdFromRefreshToken(rt.Token)
	if err != nil {
		er := errors.Wrap(err, "failed to get user id")
		s.Logger.Error(er.Error())
		return nil, er
	}

	username, email, _, err := s.Repo.GetUserByID(ctx, id)
	if err != nil {
		er := errors.Wrap(err, "failed to get user info")
		s.Logger.Error(er.Error())
		return nil, er
	}

	accessToken, err := tokens.GenerateAccessToken(id, username, email)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	s.Logger.Info("CheckRefreshToken has successfully finished")
	return &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: rt.Token,
	}, nil
}
