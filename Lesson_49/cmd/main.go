package main

import (
	"fmt"
	"log"
	"net"
	"podcast_service/config"
	pbE "podcast_service/genproto/episodes"
	pbP "podcast_service/genproto/podcasts"
	"podcast_service/service"
	"podcast_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().PODCAST_SERVICE_PORT)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	p := service.NewPodcastService(db)
	e := service.NewEpisodeService(db)

	grpcServer := grpc.NewServer()
	pbP.RegisterPodcastsServer(grpcServer, p)
	pbE.RegisterEpisodesServiceServer(grpcServer, e)

	fmt.Printf("server is listening on port %s", config.Load().PODCAST_SERVICE_PORT)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
