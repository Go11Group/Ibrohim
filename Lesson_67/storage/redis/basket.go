package redis

import (
	"context"
	pb "product-service/genproto/basket"
	"product-service/storage"

	"github.com/redis/go-redis/v9"
)

type BasketRepo struct {
	redis *redis.Client
}

func NewBasketRepo(redis *redis.Client) storage.IBasketStorage {
	return &BasketRepo{redis: redis}
}

func (b *BasketRepo) AddBasket(ctx context.Context, req *pb.NewProduct, price float32) error {
	return nil
}
