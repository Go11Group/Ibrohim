package redis

import (
	"context"
	pb "product-service/genproto/basket"
	"testing"
)

type key string

func RedisCon() *BasketRepo {
	return &BasketRepo{
		redis: ConnectDB(),
	}
}

func TestAddProduct(t *testing.T) {
	r := RedisCon()
	ctx := context.WithValue(context.Background(), key("user_id"), "1")

	err := r.AddProduct(ctx, &pb.NewProduct{
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
	ctx := context.WithValue(context.Background(), key("user_id"), "1")

	_, err := r.GetBasket(ctx)
	if err != nil {
		t.Errorf("Error occured while getting basket: %v", err)
		return
	}
}

func TestUpdateQuantity(t *testing.T) {
	r := RedisCon()
	ctx := context.WithValue(context.Background(), key("user_id"), "1")

	err := r.UpdateQuantity(ctx, &pb.Quantity{
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
	ctx := context.WithValue(context.Background(), key("user_id"), "1")

	err := r.DeleteProduct(ctx, &pb.Id{
		ProductId: "1",
	})
	if err != nil {
		t.Errorf("Error occured while deleting product: %v", err)
		return
	}
}
