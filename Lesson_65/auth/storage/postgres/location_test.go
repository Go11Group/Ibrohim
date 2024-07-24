package postgres

import (
	"auth-service/models"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Location() *LocationRepo {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	return NewLocationRepo(db)
}

func TestAddLocation(t *testing.T) {
	l := Location()
	defer l.DB.Close()

	_, err := l.Add(context.Background(), &models.NewLocation{
		UserId:     "c0a80122-0000-1000-8000-00805f9b34fb",
		Address:    "245 Crown Street",
		City:       "New York",
		State:      "NY",
		Country:    "USA",
		PostalCode: "10005",
	})
	if err != nil {
		t.Errorf("Error occurred while adding location: %v", err)
	}
}

func TestReadLocation(t *testing.T) {
	l := Location()
	defer l.DB.Close()

	loc, err := l.Read(context.Background(), "c0a80132-0000-1000-8000-00805f9b34fb")
	if err != nil {
		t.Errorf("Error occurred while reading location: %v", err)
	}

	assert.Equal(t, "c0a80132-0000-1000-8000-00805f9b34fb", loc.UserId)
}

func TestLocationRepo_Update(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer db.Close()

	// Create a LocationRepo instance
	locationRepo := &LocationRepo{DB: db}

	// Define the update data
	updateLoc := &models.NewLocation{
		UserId:     "c0a80124-0000-1000-8000-00805f9b34fb",
		Address:    "123 Updated St",
		City:       "Updated City",
		State:      "NY",
		Country:    "USA",
		PostalCode: "10001",
	}

	// Call the Update method
	resp, err := locationRepo.Update(context.Background(), updateLoc)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, updateLoc.UserId, resp.UserId)
	assert.Equal(t, updateLoc.Address, resp.Address)
	assert.Equal(t, updateLoc.City, resp.City)
	assert.Equal(t, updateLoc.State, resp.State)
	assert.Equal(t, updateLoc.Country, resp.Country)
	assert.Equal(t, updateLoc.PostalCode, resp.PostalCode)
}

func TestLocationRepo_Delete(t *testing.T) {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer db.Close()

	// Create a LocationRepo instance
	locationRepo := &LocationRepo{DB: db}

	// Define test location ID
	locationId := "c0a80132-0000-1000-8000-00805f9b34fb"

	// Call the Delete method
	err = locationRepo.Delete(context.Background(), locationId)

	// Assertions
	assert.NoError(t, err)

	// Verify the location is marked as deleted
	var deletedAt int
	err = db.QueryRow("SELECT deleted_at FROM user_locations WHERE location_id = $1", locationId).Scan(&deletedAt)
	assert.NoError(t, err)
	assert.Greater(t, deletedAt, 0, "deleted_at should be greater than 0 after deletion")
}
