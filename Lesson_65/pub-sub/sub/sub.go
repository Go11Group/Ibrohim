package main

import (
	"context"
	"fmt"
	"log"
	"stocks-management/redis"
)

func main() {
	ctx := context.Background()
	rdb := redis.ConnectDB()

	pubsub := rdb.Subscribe(ctx, "amazon", "alibaba", "aliexpress")
	defer pubsub.Close()

	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Printf("Received message from channel %s: %s\n", msg.Channel, msg.Payload)
	}
}
