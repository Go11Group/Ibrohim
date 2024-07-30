package mongodb

import (
	"context"
	pb "product-service/genproto/category"
	"product-service/storage"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepo struct {
	col *mongo.Collection
}

func NewCategoryRepo(db *mongo.Database) storage.ICategoryStorage {
	return &CategoryRepo{
		col: db.Collection("categories"),
	}
}
func (c *CategoryRepo) GetCategories(ctx context.Context, pagination *pb.Pagination) (*pb.Categories, error) {
	opts := options.Find()
	opts.SetSkip((pagination.Page - 1) * pagination.Limit)
	opts.SetLimit(pagination.Limit)

	cursor, err := c.col.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, errors.Wrap(err, "find failure")
	}
	defer cursor.Close(ctx)

	var categories *pb.Categories
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, errors.Wrap(err, "decode failure")
	}

	return categories, nil
}

func (c *CategoryRepo) CreateCategory(ctx context.Context, category *pb.Category) (*pb.Id, error) {
	result, err := c.col.InsertOne(ctx, category)
	if err != nil {
		return nil, errors.Wrap(err, "insert failure")
	}
	id := result.InsertedID.(*pb.Id)
	return id, nil
}

func (c *CategoryRepo) UpdateCategory(ctx context.Context, category *pb.Category) error {
	filter := bson.M{"_id": category.Id}
	update := bson.M{
		"$set": bson.M{
			"name":        category.Name,
			"description": category.Description,
		},
	}
	_, err := c.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "update failure")
	}
	return nil
}

func (c *CategoryRepo) DeleteCategory(ctx context.Context, id *pb.Id) error {
	filter := bson.M{"_id": id}
	_, err := c.col.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "delete failure")
	}
	return nil
}
