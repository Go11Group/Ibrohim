package mongodb

import (
	"product-service/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewRepo struct {
	col *mongo.Collection
}

func NewReviewRepo(db *mongo.Database) storage.IReviewStorage {
	return &ReviewRepo{
		col: db.Collection("reviews"),
	}
}
