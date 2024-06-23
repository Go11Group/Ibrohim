package main

import (
	"api-gateway/api"
	"fmt"
)

func main() {
	server := api.Routes()
	fmt.Println("Server is listening on port 8080...")
	panic(server.ListenAndServe())
}
