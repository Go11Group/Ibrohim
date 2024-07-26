package mongodb

import (
	"product-service/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepo struct {
	col *mongo.Collection
}

func NewCategoryRepo(db *mongo.Database) storage.ICategoryStorage {
	return &CategoryRepo{
		col: db.Collection("categories"),
	}
}
