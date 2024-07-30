package pkg

import (
	"api-gateway/config"
	pba "api-gateway/genproto/admin"
	pbb "api-gateway/genproto/basket"
	pbc "api-gateway/genproto/category"
	pbo "api-gateway/genproto/order"
	pbp "api-gateway/genproto/product"
	pbr "api-gateway/genproto/review"
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

func NewProductClient(cfg *config.Config) pbp.ProductClient {
	conn, err := grpc.NewClient(cfg.PRODUCT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbp.NewProductClient(conn)
}

func NewCategoryClient(cfg *config.Config) pbc.CategorysClient {
	conn, err := grpc.NewClient(cfg.PRODUCT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbc.NewCategorysClient(conn)
}

func NewReviewClient(cfg *config.Config) pbr.ReviewesClient {
	conn, err := grpc.NewClient(cfg.PRODUCT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbr.NewReviewesClient(conn)
}

func NewOrderClient(cfg *config.Config) pbo.OrderClient {
	conn, err := grpc.NewClient(cfg.PRODUCT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbo.NewOrderClient(conn)
}

func NewBasketClient(cfg *config.Config) pbb.BasketClient {
	conn, err := grpc.NewClient(cfg.PRODUCT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbb.NewBasketClient(conn)
}
