package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "translator/translate_service"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var words = map[string]string{"olma": "apple", "jamiyat": "society", "qor": "snow", "yugur": "run"}
var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedTranslatorServer
}

func (s *server) MustEmbedUnimplementedTranslatorServer() {}

func (s *server) Translate(ctx context.Context, in *pb.Uzbek) (*pb.English, error) {
	log.Printf("Recieved: %v", in.GetWords())
	var translatedWords []string

	for _, w := range in.GetWords() {
		translated, ok := words[w]
		if !ok {
			return nil, errors.New("failed to translate "+w)
		}
		translatedWords = append(translatedWords, translated)
	}
	
	return &pb.English{Words: translatedWords}, nil
}

func main() {
	flag.Parse()
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &server{})

	log.Printf("Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
