package main

import (
	"context"
	pb "library-service/genproto/library"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
}

func (s *server) AddBook(ctx context.Context, in *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	
}