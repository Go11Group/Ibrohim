package storage

import (
	"context"
	pbb "product-service/genproto/basket"
	"product-service/models"
)

type IStorage interface {
	Product() IProductStorage
	Category() ICategoryStorage
	Review() IReviewStorage
	Order() IOrderStorage
	Basket() IBasketStorage
	Close()
}

type IProductStorage interface{}

type ICategoryStorage interface{}

type IReviewStorage interface{}

type IOrderStorage interface{}

type IBasketStorage interface {
	AddProduct(ctx context.Context, req *pbb.NewProduct, price float32) error
	GetBasket(ctx context.Context, id *pbb.Id) (*models.Basket, error)
	UpdateQuantity(ctx context.Context, req *pbb.Quantity) error
	DeleteProduct(ctx context.Context, req *pbb.Ids) error
}
