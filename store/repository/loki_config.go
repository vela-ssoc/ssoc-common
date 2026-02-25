package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LokiConfig interface {
	Repository[model.LokiConfig]
	Enabled(ctx context.Context) (*model.LokiConfig, error)
}

func NewLokiConfig(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) LokiConfig {
	repo := NewRepository[model.LokiConfig](db, opts...)

	return &lokiConfigRepo{Repository: repo}
}

type lokiConfigRepo struct {
	Repository[model.LokiConfig]
}

func (r *lokiConfigRepo) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return r.Indexes().CreateMany(ctx, indexes, opts...)
}

func (r *lokiConfigRepo) Enabled(ctx context.Context) (*model.LokiConfig, error) {
	return r.FindOne(ctx, bson.D{{Key: "enabled", Value: true}})
}
