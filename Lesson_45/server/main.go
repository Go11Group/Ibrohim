package main

import (
	"context"
	"fmt"
	pb "library-service/genproto/library"
	"library-service/model"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var books []model.Book

type server struct {
	pb.UnimplementedLibraryServiceServer
}

func (s *server) AddBook(ctx context.Context, in *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	randomUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "failed to assign id")
	}

	newBook := model.Book{
        ID:              randomUUID.String(),
        Title:           in.Title,
        Author:          in.Author,
        PublicationYear: in.YearPublished,
        IsBorrowed:      false,
    }

	books = append(books, newBook)
	fmt.Println(books)

	return &pb.AddBookResponse{BookId: randomUUID.String()}, nil
}

func (s *server) SearchBook(ctx context.Context, in *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	res := make([]*pb.Book, 0)

	for _, v := range books {
		if in.Query == v.Title || in.Query == v.Author {
			res = append(res, &pb.Book{
				BookId:        v.ID,
				Title:         v.Title,
				Author:        v.Author,
				YearPublished: v.PublicationYear,
			})
		}
	}

	return &pb.SearchBookResponse{Books: res}, nil
}

func (s *server) BorrowBook(ctx context.Context, in *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	for _, v := range books {
		if in.BookId == v.ID {
			if v.IsBorrowed {
				return &pb.BorrowBookResponse{Status: false}, errors.New("Book already borrowed")
			} else {
				v.IsBorrowed = true
				return &pb.BorrowBookResponse{Status: true}, nil
			}
		}
	}

	return &pb.BorrowBookResponse{Status: false}, errors.New("Book not found")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLibraryServiceServer(s, &server{})

	log.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
