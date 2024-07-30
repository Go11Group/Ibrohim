package mongodb

import (
	"product-service/config"
	"product-service/storage"
	redisDB "product-service/storage/redis"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Storage struct {
	mongo *mongo.Database
	redis *redis.Client
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

	strg := Storage{
		mongo: client.Database(config.Load().MongoDB_NAME),
		redis: redisDB.ConnectDB(),
	}

	return strg, nil
}

func (s Storage) Close() {
	s.mongo.Client().Disconnect(context.Background())
	s.redis.Close()
}

func (s Storage) Product() storage.IProductStorage {
	return NewProductRepo(s.mongo)
}

func (s Storage) Category() storage.ICategoryStorage {
	return NewCategoryRepo(s.mongo)
}

func (s Storage) Review() storage.IReviewStorage {
	return NewReviewRepo(s.mongo)
}

func (s Storage) Order() storage.IOrderStorage {
	return NewOrderRepo(s.mongo)
}

func (s Storage) Basket() storage.IBasketStorage {
	return redisDB.NewBasketRepo(s.redis)
}
