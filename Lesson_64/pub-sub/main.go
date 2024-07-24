package main

import (
	"stocks-management/api"
	"stocks-management/redis"
)

func main() {
	r := api.NewRouter(redis.ConnectDB())
	r.Run(":8080")
}
