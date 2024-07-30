package mongodb

import (
	"context"
	"product-service/models"
	"product-service/storage"

	"github.com/pkg/errors"

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
func (o *OrderRepo) Purchase(ctx context.Context, userID string, basket *models.Basket) error {
	_, err := o.col.InsertOne(ctx, models.Order{
		UserId: userID,
		Items:  basket.Items,
		Sum:    basket.Sum,
	})
	if err != nil {
		return errors.Wrap(err, "insert failure")
	}
	return nil
}
