package service

import (
	"context"
	"log/slog"
	pb "product-service/genproto/product"
	"product-service/pkg/logger"
	"product-service/storage"

	"github.com/pkg/errors"
)

type ProductService struct {
	pb.UnimplementedProductServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewProductService(s storage.IStorage) *ProductService {
	return &ProductService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.NewProduct) (*pb.InsertResp, error) {
	s.logger.Info("CreateProduct is started", slog.Any("request", req))

	resp, err := s.storage.Product().Add(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to create product")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("CreateProduct is finished")
	return resp, nil
}

func (s *ProductService) GetProductById(ctx context.Context, req *pb.Id) (*pb.ProductInfo, error) {
	s.logger.Info("GetProductById is started", slog.Any("request", req))

	resp, err := s.storage.Product().Read(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get product")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("GetProductById is finished")
	return resp, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.NewData) (*pb.UpdateResp, error) {
	s.logger.Info("UpdateProduct is started", slog.Any("request", req))

	resp, err := s.storage.Product().Update(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to update product")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("UpdateProduct is finished")
	return resp, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.Id) (*pb.Void, error) {
	s.logger.Info("DeleteProduct is started", slog.Any("request", req))

	err := s.storage.Product().Delete(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to delete product")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("DeleteProduct is finished")
	return &pb.Void{}, nil
}

func (s *ProductService) FetchProducts(ctx context.Context, req *pb.Filter) (*pb.Products, error) {
	s.logger.Info("FetchProducts is started", slog.Any("request", req))

	resp, err := s.storage.Product().Fetch(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to fetch products")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("FetchProducts is finished")
	return resp, nil
}
