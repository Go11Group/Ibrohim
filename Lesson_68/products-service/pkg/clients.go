package pkg

import (
	"log"
	"product-service/config"
	pba "product-service/genproto/admin"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
