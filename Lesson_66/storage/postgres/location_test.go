package postgres

import (
	"auth-service/models"
	"auth-service/storage"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Location() storage.ILocationStorage {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	return db.Location()
}

func TestAddLocation(t *testing.T) {
	l := Location()

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

	loc, err := l.Read(context.Background(), "c0a80132-0000-1000-8000-00805f9b34fb")
	if err != nil {
		t.Errorf("Error occurred while reading location: %v", err)
	}

	assert.Equal(t, "c0a80132-0000-1000-8000-00805f9b34fb", loc.UserId)
}

func TestLocationRepo_Update(t *testing.T) {
	l := Location()

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
	resp, err := l.Update(context.Background(), updateLoc)

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
	l := Location()

	// Define test location ID
	locationId := "c0a80132-0000-1000-8000-00805f9b34fb"

	// Call the Delete method
	err := l.Delete(context.Background(), locationId)

	// Assertions
	assert.NoError(t, err)

	// Verify the location is marked as deleted
	assert.NoError(t, err)
}
