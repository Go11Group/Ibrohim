package postgres

import (
	pb "auth-service/genproto/admin"
	"auth-service/storage"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Admin() storage.IAdminStorage {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	return db.Admin()
}

func TestAdd(t *testing.T) {
	a := Admin()

	user := &pb.NewUser{
		Username: "someone",
		Email:    "somebody@somewhere.com",
		Password: "something",
		FullName: "Someone Somebody",
		Role:     "customer",
	}

	_, err := a.Add(context.Background(), user)
	if err != nil {
		t.Errorf("Error occurred while adding user: %v", err)
	}
}

func TestRead(t *testing.T) {
	a := Admin()

	exp := &pb.UserInfo{
		Id:          "c0a80122-0000-1000-8000-00805f9b34fb",
		Username:    "johndoe",
		Email:       "johndoe@example.com",
		Password:    "$2b$12$4MLlI.VpV5OeD7v8E9B5ouDl1zGyW1BWmL/6NdKUuEKkzXn3/7N96",
		FullName:    "John Doe",
		PhoneNumber: "1234567890",
		Image:       []string{"default_image.png"},
		Role:        "user",
		CreatedAt:   "2024-07-20T16:33:11.288991+05:00",
		UpdatedAt:   "2024-07-20T16:33:11.288991+05:00",
	}

	user, err := a.Read(context.Background(), &pb.ID{Id: "c0a80122-0000-1000-8000-00805f9b34fb"})
	if err != nil {
		t.Errorf("Error occurred while reading user: %v", err)
	}

	assert.Equal(t, exp, user)
}

func TestAdminRepo_Update(t *testing.T) {
	// Connect to the database
	a := Admin()

	// Define test user ID
	userID := "c0a80122-0000-1000-8000-00805f9b34fb"

	// Define the new data for update
	newData := &pb.NewData{
		Id:          userID,
		Username:    "TestJohnDoe",
		Email:       "testjohndoe@gmail.com",
		Password:    "$2b$12$4MLlI.VpV5OeD7v8E9B5ouDl1zGyW1BWmL/6NdKUuEKkzXn3/7N96",
		FullName:    "John Doe Test",
		PhoneNumber: "1234567890",
		Image:       []string{"new_image.png"},
		Role:        "admin",
	}

	// Create a context
	ctx := context.Background()

	// Call the Update method
	resp, err := a.Update(ctx, newData)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, newData.Id, resp.Id)
	assert.Equal(t, newData.Username, resp.Username)
	assert.Equal(t, newData.Email, resp.Email)
	assert.Equal(t, newData.Password, resp.Password) // Be cautious with testing password; you might need to mock or handle this differently
	assert.Equal(t, newData.FullName, resp.FullName)
	assert.Equal(t, newData.PhoneNumber, resp.PhoneNumber)
	assert.ElementsMatch(t, newData.Image, resp.Image)
	assert.Equal(t, newData.Role, resp.Role)

	assert.NoError(t, err)
}

func TestAdminRepo_Delete(t *testing.T) {
	// Connect to the database
	a := Admin()

	// Define test user ID
	userID := "c0a80131-0000-1000-8000-00805f9b34fb"

	// Define the ID to delete
	deleteID := &pb.ID{Id: userID}

	// Create a context
	ctx := context.Background()

	// Call the Delete method
	err := a.Delete(ctx, deleteID)

	// Assertions
	assert.NoError(t, err)

	assert.NoError(t, err)

	// Test deleting a non-existing user
	err = a.Delete(ctx, deleteID)
	assert.EqualError(t, err, "user not found")
}

func TestFetchUsers(t *testing.T) {
	a := Admin()

	users, err := a.FetchUsers(context.Background(), &pb.Filter{
		FullName: "John Doe",
		Role:     "user",
		Page:     1,
		Limit:    10,
	})
	if err != nil {
		t.Errorf("Error occurred while fetching users: %v", err)
	}

	assert.Equal(t, 1, len(users.Users))
}
