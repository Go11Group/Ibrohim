package mongodb

import (
	"product-service/config"
	"product-service/storage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type MongoDB struct {
	db *mongo.Database
}

func ConnectDB() (storage.IStorage, error) {
	opts := options.Client().ApplyURI(config.Load().MongoURI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return MongoDB{db: client.Database(config.Load().MongoDB_NAME)}, nil
}

func (m MongoDB) Close() {
	m.db.Client().Disconnect(context.Background())
}

func (m MongoDB) Product() storage.IProductStorage {
	return NewProductRepo(m.db)
}

func (m MongoDB) Category() storage.ICategoryStorage {
	return NewCategoryRepo(m.db)
}

func (m MongoDB) Review() storage.IReviewStorage {
	return NewReviewRepo(m.db)
}

func (m MongoDB) Order() storage.IOrderStorage {
	return NewOrderRepo(m.db)
}
