package main

import (
	"collaboration_service/config"
	pbCollab "collaboration_service/genproto/collaborations"
	pbCom "collaboration_service/genproto/comments"
	"collaboration_service/service"
	"fmt"

	"collaboration_service/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.Load().COLLABORATION_SERVICE_PORT)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pbCollab.RegisterCollaborationsServer(server, service.NewCollaborations(db))
	pbCom.RegisterCommentsServer(server, service.NewComments(db))

	fmt.Printf("server is listening on port %s", config.Load().COLLABORATION_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
