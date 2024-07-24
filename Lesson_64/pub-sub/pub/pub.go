package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"stocks-management/models"
	"stocks-management/redis"
	"time"
)

func main() {
	ctx := context.Background()
	rdb := redis.ConnectDB()

	stocks := map[string]*models.StockPrice{
		"amazon":     {Highest: -100, Lowest: 100},
		"alibaba":    {Highest: -100, Lowest: 100},
		"aliexpress": {Highest: -100, Lowest: 100},
	}

	log.Print("Publishing messages")

	for {
		time.Sleep(time.Second)
		randNum1 := rand.Intn(201) - 100
		randNum2 := rand.Intn(201) - 100
		randNum3 := rand.Intn(201) - 100

		if randNum1 > stocks["amazon"].Highest {
			stocks["amazon"].Highest = randNum1
		}
		if randNum1 < stocks["amazon"].Lowest {
			stocks["amazon"].Lowest = randNum1
		}

		if randNum2 > stocks["alibaba"].Highest {
			stocks["alibaba"].Highest = randNum2
		}
		if randNum2 < stocks["alibaba"].Lowest {
			stocks["alibaba"].Lowest = randNum2
		}

		if randNum3 > stocks["aliexpress"].Highest {
			stocks["aliexpress"].Highest = randNum3
		}
		if randNum3 < stocks["aliexpress"].Lowest {
			stocks["aliexpress"].Lowest = randNum3
		}

		if err := redis.StoreStockPrice(ctx, "amazon", stocks["amazon"]); err != nil {
			log.Fatal(err)
		}

		if err := redis.StoreStockPrice(ctx, "alibaba", stocks["alibaba"]); err != nil {
			log.Fatal(err)
		}

		if err := redis.StoreStockPrice(ctx, "aliexpress", stocks["aliexpress"]); err != nil {
			log.Fatal(err)
		}

		formattedTime := time.Now().Format("2006-01-02 15:04:05")

		err := rdb.Publish(ctx, "amazon", fmt.Sprintf("%s: %d", formattedTime, randNum1)).Err()
		if err != nil {
			log.Fatal(err)
		}

		err = rdb.Publish(ctx, "alibaba", fmt.Sprintf("%s: %d", formattedTime, randNum2)).Err()
		if err != nil {
			log.Fatal(err)
		}

		err = rdb.Publish(ctx, "aliexpress", fmt.Sprintf("%s: %d", formattedTime, randNum3)).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}
