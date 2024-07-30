package handler

import (
	"api-gateway/config"
	pba "api-gateway/genproto/admin"
	pbb "api-gateway/genproto/basket"
	pbc "api-gateway/genproto/category"
	pbo "api-gateway/genproto/order"
	pbp "api-gateway/genproto/product"
	pbr "api-gateway/genproto/review"
	pbu "api-gateway/genproto/user"
	"api-gateway/kafka/producer"
	"api-gateway/pkg"
	"api-gateway/pkg/logger"
	"log/slog"
	"time"
)

type str string

type Handler struct {
	User           pbu.UserClient
	Admin          pba.AdminClient
	Product        pbp.ProductClient
	Category       pbc.CategorysClient
	Review         pbr.ReviewesClient
	Order          pbo.OrderClient
	Basket         pbb.BasketClient
	Log            *slog.Logger
	UserIDKey      str
	ContextTimeout time.Duration
	KafkaProducer  producer.IKafkaProducer
	OrderTopic     string
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		User:           pkg.NewUserClient(cfg),
		Admin:          pkg.NewAdminClient(cfg),
		Product:        pkg.NewProductClient(cfg),
		Category:       pkg.NewCategoryClient(cfg),
		Review:         pkg.NewReviewClient(cfg),
		Order:          pkg.NewOrderClient(cfg),
		Basket:         pkg.NewBasketClient(cfg),
		Log:            logger.NewLogger(),
		UserIDKey:      str("user_id"),
		ContextTimeout: 5 * time.Second,
		KafkaProducer:  producer.NewKafkaProducer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}),
		OrderTopic:     cfg.KAFKA_TOPIC,
	}
}
