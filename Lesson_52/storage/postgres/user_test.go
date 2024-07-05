package postgres

import (
	"log"
	"reflect"
	"testing"
	pb "user_service/genproto/user"
)

func DB() *UserRepo {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	return NewUserRepo(db)
}

func TestGetUserById(t *testing.T) {
	u := DB()
	res, err := u.GetUserById("89d2d54e-005a-4ed9-b9fe-72633a90450b")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	exp := pb.User{
		Id:        "89d2d54e-005a-4ed9-b9fe-72633a90450b",
		Username:  "johndoe",
		Email:     "john.doe@email.com",
		Password:  "hashed_password_1",
		CreatedAt: "2024-07-04 10:20:18.345033",
		UpdatedAt: "",
	}

	if !reflect.DeepEqual(res, &exp) {
		t.Error(res)
		t.Errorf("error while getting user by id: %v", err)
	}

	log.Println(res, &exp)
}

func TestUpdateUser(t *testing.T) {
	u := DB()
	err := u.UpdateUser(&pb.User{Id: "4e5a4e7c-d73e-4eb4-915e-23f01646c775", Email: "amanda@mail.com"})
	if err != nil {
		t.Errorf("error while updating user: %v", err)
	}

	log.Printf("OK: %v\n", err)
}

func TestDeleteUser(t *testing.T) {
	u := DB()
	err := u.DeleteUser("4e5a4e7c-d73e-4eb4-915e-23f01646c775")
	if err != nil {
		t.Errorf("error while deleting user: %v", err)
	}

	log.Printf("OK: %v\n", err)
}

func TestGetUserProfile(t *testing.T) {
	u := DB()
	res, err := u.GetUserProfile("0aa1b84b-5a86-4283-b205-f97d83aae56f")
	exp := pb.Profile{
		UserId:   "0aa1b84b-5a86-4283-b205-f97d83aae56f",
		FullName: "Amanda Thomas",
		Bio:      "youknow",
		Role:     "listener",
		Location: "Seattle, WA",
	}

	if !reflect.DeepEqual(res, &exp) {
		t.Errorf("error while getting user profile by id: %v", err)
	}

}

func TestUpdateUserProfile(t *testing.T) {
	u := DB()
	err := u.UpdateUserProfile(&pb.Profile{
		UserId:      "4e5a4e7c-d73e-4eb4-915e-23f01646c775",
		FullName:    "Amanda Thomas James",
		Bio:         "Pop and rock music fan and playlist curator",
		Role:        "listener",
		Location:    "WA, USA",
		AvatarImage: nil,
		Website:     "",
	})
	if err != nil {
		t.Errorf("error while updating user profile: %v", err)
	}
}
