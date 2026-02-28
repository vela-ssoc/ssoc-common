package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Minion interface {
	Repository[model.Minion]
}

func NewMinion(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) Minion {
	repo := NewRepository[model.Minion](db, opts...)

	return &minionRepo{Repository: repo}
}

type minionRepo struct {
	Repository[model.Minion]
}

func (r *minionRepo) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "machine_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "inet", Value: 1}}},
		{Keys: bson.D{{Key: "tags.name", Value: 1}}},
	}

	return r.Indexes().CreateMany(ctx, indexes, opts...)
}
