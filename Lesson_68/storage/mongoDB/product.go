package mongodb

import (
	"context"
	pb "product-service/genproto/product"
	"product-service/storage"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepo struct {
	col *mongo.Collection
}

func NewProductRepo(db *mongo.Database) storage.IProductStorage {
	return &ProductRepo{
		col: db.Collection("products"),
	}
}

func (p *ProductRepo) Add(ctx context.Context, pr *pb.NewProduct) (*pb.InsertResp, error) {
	res, err := p.col.InsertOne(ctx, pr)
	if err != nil {
		return nil, errors.Wrap(err, "insertion failure")
	}

	return &pb.InsertResp{
		Id:        res.InsertedID.(string),
		CreatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (p *ProductRepo) Read(ctx context.Context, id string) (*pb.Product, error) {
	var pr pb.Product
	filter := bson.M{"_id": id}

	err := p.col.FindOne(ctx, filter).Decode(&pr)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("product not found")
		}
		return nil, errors.Wrap(err, "decode failure")
	}

	return &pr, nil
}

func (p *ProductRepo) Update(ctx context.Context, pr *pb.NewData) (*pb.UpdateResp, error) {
	_, err := p.col.UpdateOne(ctx, bson.M{"_id": pr.Id}, bson.M{"$set": pr})
	if err != nil {
		return nil, errors.Wrap(err, "update failure")
	}

	return &pb.UpdateResp{
		Id:        pr.Id,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (p *ProductRepo) Delete(ctx context.Context, id string) error {
	res, err := p.col.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrap(err, "deletion failure")
	}

	if res.DeletedCount < 1 {
		return errors.New("product not found")
	}

	return nil
}

func (p *ProductRepo) Fetch(ctx context.Context, f *pb.Filter) (*pb.Products, error) {
	filter := bson.M{}
	opts := options.Find()

	if f.Name != "" {
		filter["name"] = bson.M{"$regex": f.Name, "$options": "i"}
	}
	if f.Category != "" {
		filter["category"] = bson.M{"$regex": f.Category, "$options": "i"}
	}
	if f.CommentCount > 0 {
		filter["comment_count"] = bson.M{"$gte": f.CommentCount}
	}
	if f.Rating > 0 {
		filter["rating"] = bson.M{"$gte": f.Rating}
	}
	if f.Discount {
		filter["discount.status"] = true
	}

	if f.MostPurchased {
		opts.SetSort(bson.M{"purchase_count": -1})
	} else if f.MostCommented {
		opts.SetSort(bson.M{"comment_count": -1})
	} else if f.MostRecent {
		opts.SetSort(bson.M{"created_at": -1})
	} else if f.Cheapest {
		opts.SetSort(bson.M{"price": 1})
	} else if f.MostExpensive {
		opts.SetSort(bson.M{"price": -1})
	}

	if f.Page > 0 && f.Limit > 0 {
        opts.SetSkip(int64(f.Page * f.Limit - f.Limit))
        opts.SetLimit(int64(f.Limit))
    }

	cur, err := p.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "retrieval failure")
	}
	defer cur.Close(ctx)

	var products []*pb.Product
	for cur.Next(ctx) {
		var pr pb.Product
		if err := cur.Decode(&pr); err != nil {
			return nil, errors.Wrap(err, "decode failure")
		}
		products = append(products, &pr)
	}

	if err := cur.Err(); err != nil {
        return nil, errors.Wrap(err, "cursor error")
    }

	return &pb.Products{
		Products: products,
		Page:     f.Page,
		Limit:    f.Limit,
	}, nil
}

func (p *ProductRepo) GetName(ctx context.Context, id string) (string, error) {
	var res struct {
		Name string `bson:"name"`
	}
	filter := bson.M{"_id": id}
	opts := options.FindOne().SetProjection(bson.M{"name": 1})

	err := p.col.FindOne(ctx, filter, opts).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("product not found")
		}
		return "", errors.Wrap(err, "decode failure")
	}

	return res.Name, nil
}

func (p *ProductRepo) GetPrice(ctx context.Context, id string) (float32, error) {
	var res struct {
		Price float32 `bson:"price"`
	}
	filter := bson.M{"_id": id}
	opts := options.FindOne().SetProjection(bson.M{"price": 1})

	err := p.col.FindOne(ctx, filter, opts).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return -1, errors.New("product not found")
		}
		return -1, errors.Wrap(err, "decode failure")
	}

	return res.Price, nil
}

func (p *ProductRepo) Validate(ctx context.Context, id string) error {
	res, err := p.col.CountDocuments(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrap(err, "validation failure")
	}

	if res < 1 {
		return errors.New("no such product")
	}

	return nil
}
