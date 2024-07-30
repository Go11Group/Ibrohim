package storage

import (
	"context"
	pbb "product-service/genproto/basket"
	pbp "product-service/genproto/product"
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

type IProductStorage interface {
	Add(ctx context.Context, pr *pbp.NewProduct) (*pbp.InsertResp, error)
	Read(ctx context.Context, id string) (*pbp.ProductInfo, error)
	Update(ctx context.Context, pr *pbp.NewData) (*pbp.UpdateResp, error)
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, filter *pbp.Filter) (*pbp.Products, error)
	GetName(ctx context.Context, id string) (string, error)
	GetPrice(ctx context.Context, id string) (float32, error)
	Validate(ctx context.Context, id string) error
}

type ICategoryStorage interface{}

type IReviewStorage interface{}

type IOrderStorage interface {
	Purchase(ctx context.Context, userID string, basket *models.Basket) error
}

type IBasketStorage interface {
	AddProduct(ctx context.Context, req *pbb.NewProduct, price float32) error
	GetBasket(ctx context.Context) (*models.Basket, error)
	UpdateQuantity(ctx context.Context, req *pbb.Quantity) error
	DeleteProduct(ctx context.Context, req *pbb.Id) error
}
