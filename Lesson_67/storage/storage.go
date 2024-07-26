package storage

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

type IBasketStorage interface{}
