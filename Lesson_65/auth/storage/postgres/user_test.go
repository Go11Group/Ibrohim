package postgres

import (
	pb "auth-service/genproto/user"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserRepo_GetProfile(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer db.Close()

	// Create a UserRepo instance
	userRepo := &UserRepo{DB: db}

	// Define test user ID
	userID := "c0a80123-0000-1000-8000-00805f9b34fb"

	// Define expected profile data
	expectedProfile := &pb.Profile{
		Id:          userID,
		FullName:    "Jane Smith",
		Username:    "janesmith",
		Email:       "janesmith@example.com",
		PhoneNumber: "0987654321",
		Image:       []string{"default_image.png"},
		Role:        "user",
		CreatedAt:   "2024-07-20T15:45:02.381261+05:00",
		UpdatedAt:   "2024-07-20T15:45:02.381261+05:00",
	}

	// Create a context with the user ID
	ctx := context.WithValue(context.Background(), "user_id", userID)

	// Call the GetProfile method
	profile, err := userRepo.GetProfile(ctx)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedProfile, profile)
}

func TestUserRepo_UpdateProfile(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer db.Close()

	// Create a UserRepo instance
	userRepo := &UserRepo{DB: db}

	// Define test user ID
	userID := "c0a80124-0000-1000-8000-00805f9b34fb"

	// Define the new data to update
	newData := &pb.NewData{
		FullName:    "Alice Johnson Test",
		Username:    "alicej",
		Email:       "alicej@example.com",
		PhoneNumber: "5551234567",
		Image:       []string{"new_image.png"},
	}

	// Define the expected response data after update
	expectedResp := &pb.UpdateResp{
		Id:          userID,
		FullName:    newData.FullName,
		Username:    newData.Username,
		Email:       newData.Email,
		PhoneNumber: newData.PhoneNumber,
		Image:       newData.Image,
		UpdatedAt:   time.Now().String(), // The actual updated_at value should match the value in the database
	}

	// Create a context with the user ID
	ctx := context.WithValue(context.Background(), "user_id", userID)

	// Call the UpdateProfile method
	resp, err := userRepo.UpdateProfile(ctx, newData)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.Id, resp.Id)
	assert.Equal(t, expectedResp.Username, resp.Username)
	assert.Equal(t, expectedResp.FullName, resp.FullName)
	assert.Equal(t, expectedResp.Email, resp.Email)
	assert.Equal(t, expectedResp.PhoneNumber, resp.PhoneNumber)
	assert.ElementsMatch(t, expectedResp.Image, resp.Image)
}

func TestUserRepo_DeleteProfile(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer db.Close()

	// Create a UserRepo instance
	userRepo := &UserRepo{DB: db}

	// Define test user ID
	userID := "c0a80125-0000-1000-8000-00805f9b34fb"

	// Create a context with the user ID
	ctx := context.WithValue(context.Background(), "user_id", userID)

	// Call the DeleteProfile method
	err = userRepo.DeleteProfile(ctx)

	// Assertions
	assert.NoError(t, err)

	// Verify the user is marked as deleted
	var deletedAt int
	err = db.QueryRow("SELECT deleted_at FROM users WHERE id = $1", userID).Scan(&deletedAt)
	assert.NoError(t, err)
	assert.Greater(t, deletedAt, 0, "deleted_at should be greater than 0 after deletion")
}
