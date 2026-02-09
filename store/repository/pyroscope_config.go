package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type PyroscopeConfig interface {
	Repository[model.PyroscopeConfig]
	Enabled(ctx context.Context) (*model.PyroscopeConfig, error)
}

func NewPyroscopeConfig(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) PyroscopeConfig {
	repo := NewRepository[model.PyroscopeConfig](db, opts...)

	return &pyroscopeConfigRepo{Repository: repo}
}

type pyroscopeConfigRepo struct {
	Repository[model.PyroscopeConfig]
}

func (r *pyroscopeConfigRepo) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return r.Indexes().CreateMany(ctx, indexes, opts...)
}

func (r *pyroscopeConfigRepo) Enabled(ctx context.Context) (*model.PyroscopeConfig, error) {
	return r.FindOne(ctx, bson.D{{Key: "enabled", Value: true}})
}
