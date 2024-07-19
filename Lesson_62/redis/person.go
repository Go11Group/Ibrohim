package redis

import (
	"context"
	"encoding/json"
	"redis-crud/models"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type PersonRepo struct {
	DB *redis.Client
}

func NewPersonRepo(rdb *redis.Client) *PersonRepo {
	return &PersonRepo{DB: rdb}
}

func (p *PersonRepo) Add(ctx context.Context, data *models.PersonInfo) (*models.Person, error) {
	pn := models.Person{
		ID:        uuid.NewString(),
		Name:      data.Name,
		Age:       data.Age,
		IsMarried: data.IsMarried,
	}

	if err := p.DB.Set(ctx, pn.ID, data, 0).Err(); err != nil {
		return nil, errors.Wrap(err, "failed to add person")
	}

	return &pn, nil
}

func (p *PersonRepo) Read(ctx context.Context, id string) (*models.PersonInfo, error) {
	res := p.DB.Get(ctx, id)
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "failed to read person")
	}

	result, err := res.Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get result from redis")
	}

	var pn models.PersonInfo
	if err := json.Unmarshal([]byte(result), &pn); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal person info")
	}

	return &pn, nil
}
