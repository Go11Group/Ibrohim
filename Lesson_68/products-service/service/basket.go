package service

import (
	"context"
	"log/slog"
	pba "product-service/genproto/admin"
	pb "product-service/genproto/basket"
	"product-service/pkg/logger"
	"product-service/storage"

	"github.com/pkg/errors"
)

type BasketService struct {
	pb.UnimplementedBasketServer
	storage storage.IStorage
	logger  *slog.Logger
	user    pba.AdminClient
}

func NewBasketService(s storage.IStorage, client pba.AdminClient) *BasketService {
	return &BasketService{
		storage: s,
		logger:  logger.NewLogger(),
		user:    client,
	}
}

func (s *BasketService) AddProduct(ctx context.Context, req *pb.NewProduct) (*pb.Void, error) {
	s.logger.Info("AddProduct is started", slog.Any("request", req))

	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		er := errors.New("failed to get user id")
		s.logger.Error(er.Error())
		return nil, er
	}

	status, err := s.user.ValidateUser(ctx, &pba.ID{Id: userId})
	if err != nil || !status.Status {
		er := errors.Wrap(err, "failed to validate user")
		s.logger.Error(er.Error())
		return nil, er
	}

	if err := s.storage.Product().Validate(ctx, req.ProductId); err != nil {
		er := errors.Wrap(err, "failed to validate product")
		s.logger.Error(er.Error())
		return nil, er
	}

	price, err := s.storage.Product().GetPrice(ctx, req.ProductId)
	if err != nil {
		er := errors.Wrap(err, "failed to get product price")
		s.logger.Error(er.Error())
		return nil, er
	}

	err = s.storage.Basket().AddProduct(ctx, req, price)
	if err != nil {
		er := errors.Wrap(err, "failed to add product")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("AddProduct is finished")
	return &pb.Void{}, nil
}

func (s *BasketService) GetProducts(ctx context.Context, req *pb.Void) (*pb.Products, error) {
	s.logger.Info("GetProducts is started", slog.Any("request", req))

	_, ok := ctx.Value("user_id").(string)
	if !ok {
		er := errors.New("failed to get user id")
		s.logger.Error(er.Error())
		return nil, er
	}

	basket, err := s.storage.Basket().GetBasket(ctx)
	if err != nil {
		er := errors.Wrap(err, "failed to get basket products")
		s.logger.Error(er.Error())
		return nil, er
	}

	items := make([]*pb.Product, 0, len(basket.Items))

	for _, item := range basket.Items {
		name, err := s.storage.Product().GetName(ctx, item.ProductId)
		if err != nil {
			er := errors.Wrap(err, "failed to get product name")
			s.logger.Error(er.Error())
			return nil, er
		}

		items = append(items, &pb.Product{
			Id:       item.ProductId,
			Name:     name,
			Price:    item.Price,
			Quantity: item.Quantity,
		})
	}

	s.logger.Info("GetProducts is finished")
	return &pb.Products{Items: items}, nil
}

func (s *BasketService) UpdateProduct(ctx context.Context, req *pb.Quantity) (*pb.Void, error) {
	s.logger.Info("UpdateProduct is started", slog.Any("request", req))

	_, ok := ctx.Value("user_id").(string)
	if !ok {
		er := errors.New("failed to get user id")
		s.logger.Error(er.Error())
		return nil, er
	}

	err := s.storage.Basket().UpdateQuantity(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to update product quantity")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("UpdateProduct is finished")
	return &pb.Void{}, nil
}

func (s *BasketService) RemoveProduct(ctx context.Context, req *pb.Id) (*pb.Void, error) {
	s.logger.Info("RemoveProduct is started", slog.Any("request", req))

	_, ok := ctx.Value("user_id").(string)
	if !ok {
		er := errors.New("failed to get user id")
		s.logger.Error(er.Error())
		return nil, er
	}

	err := s.storage.Basket().DeleteProduct(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to remove product")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("RemoveProduct is finished")
	return &pb.Void{}, nil
}
