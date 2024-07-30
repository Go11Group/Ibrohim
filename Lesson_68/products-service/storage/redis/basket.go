package redis

import (
	"context"
	"encoding/json"
	pb "product-service/genproto/basket"
	"product-service/models"
	"product-service/storage"
	"strconv"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type BasketRepo struct {
	redis *redis.Client
}

func NewBasketRepo(redis *redis.Client) storage.IBasketStorage {
	return &BasketRepo{redis: redis}
}

func (b *BasketRepo) AddProduct(ctx context.Context, req *pb.NewProduct, price float32) error {
	basketKey := "basket:" + ctx.Value("user_id").(string)

	basket, err := b.GetBasket(ctx)
	if err != nil {
		return errors.Wrap(err, "basket retrieval failure")
	}

	basket.Items = append(basket.Items, models.Product{
		ProductId: req.ProductId,
		Price:     price,
		Quantity:  req.Quantity,
	})
	newSum := float32(basket.Sum) + price*float32(req.Quantity)

	updatedItemsStr, err := json.Marshal(basket.Items)
	if err != nil {
		return errors.Wrap(err, "items marshaling failure")
	}

	_, err = b.redis.HSet(ctx, basketKey, "sum", newSum).Result()
	if err != nil {
		return errors.Wrap(err, "basket sum update failure")
	}

	_, err = b.redis.HSet(ctx, basketKey, "items", updatedItemsStr).Result()
	if err != nil {
		return errors.Wrap(err, "basket items update failure")
	}

	return nil
}

func (b *BasketRepo) GetBasket(ctx context.Context) (*models.Basket, error) {
	basketKey := "basket:" + ctx.Value("user_id").(string)

	sumStr, err := b.redis.HGet(ctx, basketKey, "sum").Result()
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, "sum retrieval failure")
	}

	itemsStr, err := b.redis.HGet(ctx, basketKey, "items").Result()
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, "items retrieval failure")
	}

	var currentSum float64
	if err == redis.Nil {
		currentSum = 0
		itemsStr = "[]"
	} else {
		currentSum, err = strconv.ParseFloat(sumStr, 32)
		if err != nil {
			return nil, errors.Wrap(err, "sum conversion failure")
		}
	}

	var items []models.Product
	if err := json.Unmarshal([]byte(itemsStr), &items); err != nil {
		return nil, errors.Wrap(err, "items unmarshaling failure")
	}

	return &models.Basket{
		Items: items,
		Sum:   float32(currentSum),
	}, nil
}

func (b *BasketRepo) UpdateQuantity(ctx context.Context, req *pb.Quantity) error {
	basketKey := "basket:" + ctx.Value("user_id").(string)

	basket, err := b.GetBasket(ctx)
	if err != nil {
		return errors.Wrap(err, "basket retrieval failure")
	}

	for i, item := range basket.Items {
		if item.ProductId == req.ProductId {
			basket.Items[i].Quantity = req.Quantity
		}
	}

	updatedItemsStr, err := json.Marshal(basket.Items)
	if err != nil {
		return errors.Wrap(err, "items marshaling failure")
	}

	_, err = b.redis.HSet(ctx, basketKey, "items", updatedItemsStr).Result()
	if err != nil {
		return errors.Wrap(err, "basket items update failure")
	}

	return nil
}

func (b *BasketRepo) DeleteProduct(ctx context.Context, req *pb.Id) error {
	basketKey := "basket:" + ctx.Value("user_id").(string)
	basket, err := b.GetBasket(ctx)
	if err != nil {
		return errors.Wrap(err, "basket retrieval failure")
	}

	for i, item := range basket.Items {
		if item.ProductId == req.ProductId {
			basket.Items = append(basket.Items[:i], basket.Items[i+1:]...)
		}
	}

	updatedItemsStr, err := json.Marshal(basket.Items)
	if err != nil {
		return errors.Wrap(err, "items marshaling failure")
	}

	_, err = b.redis.HSet(ctx, basketKey, "items", updatedItemsStr).Result()
	if err != nil {
		return errors.Wrap(err, "basket items update failure")
	}

	return nil
}
