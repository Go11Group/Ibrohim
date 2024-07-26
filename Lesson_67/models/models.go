package models

type Product struct {
	ProductId string  `json:"product_id"`
	Price     float32 `json:"price"`
	Quantity  int64   `json:"quantity"`
}

type Basket struct {
	Items []Product `json:"items"`
	Sum   float32   `json:"sum"`
}
