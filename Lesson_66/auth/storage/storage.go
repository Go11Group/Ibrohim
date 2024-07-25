package storage

import (
	pba "auth-service/genproto/admin"
	pbu "auth-service/genproto/user"
	"auth-service/models"
	"context"
)

type IStorage interface {
	Admin() IAdminStorage
	User() IUserStorage
	Location() ILocationStorage
	Token() ITokenStorage
	Close()
}

type IAdminStorage interface {
	Add(ctx context.Context, u *pba.NewUser) (*pba.NewUserResp, error)
	Read(ctx context.Context, id *pba.ID) (*pba.UserInfo, error)
	Update(ctx context.Context, data *pba.NewData) (*pba.NewDataResp, error)
	Delete(ctx context.Context, id *pba.ID) error
	FetchUsers(ctx context.Context, f *pba.Filter) (*pba.Users, error)
}

type IUserStorage interface {
	GetProfile(ctx context.Context) (*pbu.Profile, error)
	UpdateProfile(ctx context.Context, req *pbu.NewData) (*pbu.UpdateResp, error)
	DeleteProfile(ctx context.Context) error
	GetUserByID(ctx context.Context, id string) (string, string, string, error)
	GetUserByEmail(ctx context.Context, email string) (string, string, string, error)
}

type ILocationStorage interface {
	Add(ctx context.Context, newLoc *models.NewLocation) (*models.LocationDetails, error)
	Read(ctx context.Context, id string) (*models.LocationDetails, error)
	Update(ctx context.Context, newLoc *models.NewLocation) (*models.UpdateResp, error)
	Delete(ctx context.Context, id string) error
}

type ITokenStorage interface {
	Store(ctx context.Context, token *models.RefreshTokenDetails) error
	Delete(ctx context.Context, UserId string) error
}
