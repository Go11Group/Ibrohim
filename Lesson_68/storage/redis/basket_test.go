package redis

import (
	"context"
	pb "product-service/genproto/basket"
	"testing"
)

func RedisCon() *BasketRepo {
	return &BasketRepo{
		redis: ConnectDB(),
	}
}

func TestAddProduct(t *testing.T) {
	r := RedisCon()

	err := r.AddProduct(context.Background(), &pb.NewProduct{
		UserId:    "1",
		ProductId: "1",
		Quantity:  1,
	}, 10)
	if err != nil {
		t.Errorf("Error occured while adding product: %v", err)
		return
	}
}

func TestGetBasket(t *testing.T) {
	r := RedisCon()

	_, err := r.GetBasket(context.Background(), &pb.Id{UserId: "1"})
	if err != nil {
		t.Errorf("Error occured while getting basket: %v", err)
		return
	}
}

func TestUpdateQuantity(t *testing.T) {
	r := RedisCon()

	err := r.UpdateQuantity(context.Background(), &pb.Quantity{
		UserId:    "1",
		ProductId: "1",
		Quantity:  5,
	})
	if err != nil {
		t.Errorf("Error occured while updating quantity: %v", err)
		return
	}
}

func TestDeleteProduct(t *testing.T) {
	r := RedisCon()

	err := r.DeleteProduct(context.Background(), &pb.Ids{
		UserId:    "1",
		ProductId: "1",
	})
	if err != nil {
		t.Errorf("Error occured while deleting product: %v", err)
		return
	}
}
