package main

import (
	"context"
	"log"
	"net"
	"product-service/config"
	pbb "product-service/genproto/basket"
	pbo "product-service/genproto/order"
	pbp "product-service/genproto/product"
	"product-service/kafka/consumer"
	"product-service/pkg"
	"product-service/service"
	mongodb "product-service/storage/mongoDB"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	db, err := mongodb.ConnectDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", cfg.PRODUCT_SERVICE_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	p := service.NewProductService(db)
	b := service.NewBasketService(db, pkg.NewAdminClient(cfg))
	o := service.NewOrderService(db)
	server := grpc.NewServer()
	pbp.RegisterProductServer(server, p)
	pbb.RegisterBasketServer(server, b)
	pbo.RegisterOrderServer(server, o)

	consumer := consumer.NewKafkaConsumer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}, cfg.KAFKA_TOPIC)
	go func() {
		consumer.Consume(func(message []byte) {
			log.Printf("Received message: %s", string(message))

			_, err = o.Purchase(context.Background(), &pbo.Msg{
				UserId: message,
			})

			if err != nil {
				log.Printf("Error occured while purchasing: %s", err)
			}
		})
	}()

	log.Printf("Service is listening on port %s...\n", cfg.PRODUCT_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error while serving product service: %s", err)
	}
}
