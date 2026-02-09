package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type VictoriaMetricsConfig interface {
	Repository[model.VictoriaMetricsConfig]
	Enabled(ctx context.Context) (*model.VictoriaMetricsConfig, error)
}

func NewVictoriaMetricsConfig(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) VictoriaMetricsConfig {
	repo := NewRepository[model.VictoriaMetricsConfig](db, opts...)

	return &victoriaMetricsConfigRepo{Repository: repo}
}

type victoriaMetricsConfigRepo struct {
	Repository[model.VictoriaMetricsConfig]
}

func (r *victoriaMetricsConfigRepo) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return r.Indexes().CreateMany(ctx, indexes, opts...)
}

func (r *victoriaMetricsConfigRepo) Enabled(ctx context.Context) (*model.VictoriaMetricsConfig, error) {
	return r.FindOne(ctx, bson.D{{Key: "enabled", Value: true}})
}
