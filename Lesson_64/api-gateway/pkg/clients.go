package pkg

import (
	"api-gateway/config"
	pba "api-gateway/genproto/admin"
	pbu "api-gateway/genproto/user"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(cfg *config.Config) pbu.UserClient {
	conn, err := grpc.NewClient(cfg.AUTH_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbu.NewUserClient(conn)
}

func NewAdminClient(cfg *config.Config) pba.AdminClient {
	conn, err := grpc.NewClient(cfg.AUTH_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pba.NewAdminClient(conn)
}
