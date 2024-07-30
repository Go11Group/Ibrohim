package models

type Product struct {
	ProductId string  `json:"product_id"`
	Price     float32 `json:"price"`
	Quantity  int64   `json:"quantity"`
}

type Basket struct {
	Items []Product `redis:"items"`
	Sum   float32   `redis:"sum"`
}

type Order struct {
	UserId string    `bson:"user_id"`
	Items  []Product `bson:"items"`
	Sum    float32   `bson:"sum"`
}
type Review struct {
	ProductId string `bson:"product_id"`
	UserId    string `bson:"user_id"`
	Rating    int32  `bson:"rating"`
	Comment   string `bson:"comment"`
}
