package mongodb

import (
	"context"
	pb "product-service/genproto/review"
	"product-service/models"
	"product-service/storage"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewRepo struct {
	col *mongo.Collection
}

func NewReviewRepo(db *mongo.Database) storage.IReviewStorage {
	return &ReviewRepo{
		col: db.Collection("reviews"),
	}
}

func (r *ReviewRepo) CreateReview(ctx context.Context, review *pb.NewData) (*pb.Id, error) {
	UserId := ctx.Value("user_id").(string)

	result, err := r.col.InsertOne(ctx, models.Review{ProductId: review.ProductId, UserId: UserId, Rating: review.Rating, Comment: review.Comment})
	if err != nil {
		return nil, errors.Wrap(err, "insert failure")
	}
	id := result.InsertedID.(*pb.Id)
	return id, nil
}

func (r *ReviewRepo) UpdateReview(ctx context.Context, review *pb.UReview) (*pb.Review, error) {
	filter := bson.M{"_id": review.Id}
	update := bson.M{
		"$set": bson.M{
			"rating":  review.Rating,
			"comment": review.Comment,
		},
	}
	_, err := r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, errors.Wrap(err, "update failure")
	}
	var result *pb.Review
	err = r.col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, errors.Wrap(err, "find failure")
	}
	return result, nil
}

func (r *ReviewRepo) DeleteReview(ctx context.Context, id *pb.Id) error {
	filter := bson.M{"_id": id}
	_, err := r.col.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "delete failure")
	}
	return nil
}

func (r *ReviewRepo) FetchReviews(ctx context.Context, pagination *pb.Pagination) (*pb.Reviews, error) {
	opts := options.Find()
	opts.SetSkip((pagination.Page - 1) * pagination.Limit)
	opts.SetLimit(pagination.Limit)

	cursor, err := r.col.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, errors.Wrap(err, "find failure")
	}
	defer cursor.Close(ctx)

	var reviews pb.Reviews
	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, errors.Wrap(err, "decode failure")
	}

	return &reviews, nil
}
