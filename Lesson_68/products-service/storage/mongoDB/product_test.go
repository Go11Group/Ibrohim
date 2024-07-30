package mongodb

import (
	"product-service/config"
	pb "product-service/genproto/product"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func ProductCon() *ProductRepo {
	opts := options.Client().ApplyURI(config.Load().MongoURI)
	client, _ := mongo.Connect(context.Background(), opts)

	return &ProductRepo{
		col: client.Database(config.Load().MongoDB_NAME).Collection("products"),
	}
}

func TestAddProduct(t *testing.T) {
	p := ProductCon()
	image := &pb.Image{ImageUrl: "test"}

	_, err := p.Add(context.Background(), &pb.NewProduct{
		Name:        "test",
		Description: "test",
		Category:    "test",
		Price:       1,
		Stock:       1,
		Discount: &pb.Discount{
			DiscountPrice: 1,
			Status:        true,
		},
		Images: []*pb.Image{image},
	})

	if err != nil {
		t.Errorf("Error occured while adding product: %v", err)
		return
	}
}

func TestReadProduct(t *testing.T) {
	p := ProductCon()

	_, err := p.Read(context.Background(), "66a77edb6d734dde15206045")
	if err != nil {
		t.Errorf("Error occured while reading product: %v", err)
		return
	}
}

func TestUpdateProduct(t *testing.T) {
	p := ProductCon()

	_, err := p.Update(context.Background(), &pb.NewData{
		Id:    "66a77edb6d734dde15206045",
		Name:  "test",
		Price: 1,
		Stock: 1,
	})
	if err != nil {
		t.Errorf("Error occured while updating product: %v", err)
		return
	}
}

func TestDeleteProduct(t *testing.T) {
	p := ProductCon()

	err := p.Delete(context.Background(), "66a77edb6d734dde15206045")
	if err != nil {
		t.Errorf("Error occured while deleting product: %v", err)
		return
	}
}

func TestFetchProduct(t *testing.T) {
	p := ProductCon()

	_, err := p.Fetch(context.Background(), &pb.Filter{})
	if err != nil {
		t.Errorf("Error occured while fetching product: %v", err)
		return
	}
}
