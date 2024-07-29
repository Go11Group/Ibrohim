package mongodb

import (
	"product-service/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepo struct {
	col *mongo.Collection
}

func NewOrderRepo(db *mongo.Database) storage.IOrderStorage {
	return &OrderRepo{
		col: db.Collection("orders"),
	}
}
