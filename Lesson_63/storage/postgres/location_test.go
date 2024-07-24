package postgres

import (
	pb "auth-service/genproto/location"
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

	_, err := l.Add(context.Background(), &pb.NewLocation{
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

	loc, err := l.Read(context.Background(), &pb.ID{Id: "c0a80132-0000-1000-8000-00805f9b34fb"})
	if err != nil {
		t.Errorf("Error occurred while reading location: %v", err)
	}

	assert.Equal(t, "c0a80132-0000-1000-8000-00805f9b34fb", loc.Id)
}
