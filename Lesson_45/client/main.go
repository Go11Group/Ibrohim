package main

import (
	"context"
	"fmt"
	pb "library-service/genproto/library"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to the address: %v", err)
	}
	defer conn.Close()

	c := pb.NewLibraryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// AddBook(c, ctx, "Forever", "Jonathon Clark", 2021)
	// AddBook(c, ctx, "Just to be", "Anna Smith", 2017)

	// SearchBook(c, ctx, "Forever")
	// SearchBook(c, ctx, "Anna Smith")

	BorrowBook(c, ctx, "9ae3302c-88f7-4c58-a218-d4fd670f7f92", "456")
}

func AddBook(c pb.LibraryServiceClient, ctx context.Context, t, a string, yp int32) {
	req := &pb.AddBookRequest{Title: t, Author: a, YearPublished: yp}
	resp, err := c.AddBook(ctx, req)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	fmt.Println(resp)
}

func SearchBook(c pb.LibraryServiceClient, ctx context.Context, q string) {
	req := &pb.SearchBookRequest{Query: q}
	resp, err := c.SearchBook(ctx, req)
	if err != nil {
		log.Fatalf("could not find: %v", err)
	}
	fmt.Println(resp)
}

func BorrowBook(c pb.LibraryServiceClient, ctx context.Context, b, u string) {
	req := &pb.BorrowBookRequest{BookId: b, UserId: u}
	resp, err := c.BorrowBook(ctx, req)
	if err != nil {
		log.Fatalf("could not borrow: %v", err)
	}
	fmt.Println(resp)
}
