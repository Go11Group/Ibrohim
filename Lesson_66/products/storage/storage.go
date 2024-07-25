package storage

type IStorage interface {
	Product() IProductStorage
	Category() ICategoryStorage
	Review() IReviewStorage
	Order() IOrderStorage
	Close()
}

type IProductStorage interface{}

type ICategoryStorage interface{}

type IReviewStorage interface{}

type IOrderStorage interface{}
