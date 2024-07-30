package service

import (
	"context"
	"log/slog"
	pb "product-service/genproto/order"
	"product-service/pkg/logger"
	"product-service/storage"

	"github.com/pkg/errors"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewOrderService(s storage.IStorage) *OrderService {
	return &OrderService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (s *OrderService) Purchase(ctx context.Context, req *pb.Msg) (*pb.Void, error) {
	s.logger.Info("Purchase is started", slog.Any("request", req))

	basket, err := s.storage.Basket().GetBasket(ctx)
	if err != nil {
		er := errors.Wrap(err, "failed to get basket")
		s.logger.Error(er.Error())
		return nil, er
	}

	err = s.storage.Order().Purchase(ctx, string(req.UserId), basket)
	if err != nil {
		er := errors.Wrap(err, "failed to purchase")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("Purchase is finished")
	return &pb.Void{}, nil
}
