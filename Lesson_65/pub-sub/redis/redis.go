package redis

import (
	"context"
	"encoding/json"
	"stocks-management/models"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func ConnectDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func StoreStockPrice(ctx context.Context, stockName string, stockPrice *models.StockPrice) error {
	rdb := ConnectDB()

	bytes, err := json.Marshal(stockPrice)
	if err != nil {
		return errors.Wrap(err, "failed to marshal stock price")
	}

	err = rdb.HSet(ctx, "stocks", stockName, bytes).Err()
	if err != nil {
		return errors.Wrap(err, "failed to store stock price")
	}

	return nil
}

func GetStockPrice(ctx context.Context, stockName string) (*models.StockPrice, error) {
	rdb := ConnectDB()

	res, err := rdb.HGet(ctx, "stocks", stockName).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stock price")
	}

	var stockPrice models.StockPrice
	err = json.Unmarshal([]byte(res), &stockPrice)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal stock price")
	}

	return &stockPrice, nil
}
