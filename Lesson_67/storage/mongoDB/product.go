package mongodb

import (
	"context"
	"product-service/storage"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepo struct {
	col *mongo.Collection
}

func NewProductRepo(db *mongo.Database) storage.IProductStorage {
	return &ProductRepo{
		col: db.Collection("products"),
	}
}

func (p *ProductRepo) GetPrice(ctx context.Context, id string) (float32, error) {
	var res struct {
		Price float32 `bson:"price"`
	}
	filter := bson.M{"_id": id}
	opts := bson.M{"price": 1}

	err := p.col.FindOne(ctx, filter, options.FindOne().SetProjection(opts)).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, errors.New("product not found")
		}
		return 0, errors.Wrap(err, "decode failure")
	}

	return res.Price, nil
}
