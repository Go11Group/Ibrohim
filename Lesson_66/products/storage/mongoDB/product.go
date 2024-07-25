package mongodb

import (
	"product-service/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepo struct {
	col *mongo.Collection
}

func NewProductRepo(db *mongo.Database) storage.IProductStorage {
	return &ProductRepo{
		col: db.Collection("products"),
	}
}
