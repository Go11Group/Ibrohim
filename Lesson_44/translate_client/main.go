package main

import (
	"context"
	"flag"
	"log"
	"time"
	pb "translator/translate_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to the address: %v", err)
	}
	defer conn.Close()

	c := pb.NewTranslatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	uzbekWords := &pb.Uzbek{
		Words: []string{"olma", "jamiyat", "qor", "yugur"},
	}

	resp, err := c.Translate(ctx, uzbekWords)
	if err != nil {
		log.Fatalf("could not translate: %v", err)
	}

	log.Printf("English words: %v", resp.GetWords())
}
