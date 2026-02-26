package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Certificate interface {
	Repository[model.Certificate]
	Enables(ctx context.Context) ([]*model.Certificate, error)
}

func NewCertificate(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) Certificate {
	repo := NewRepository[model.Certificate](db, opts...)

	return &certificateRepo{Repository: repo}
}

type certificateRepo struct {
	Repository[model.Certificate]
}

func (r *certificateRepo) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "certificate_sha256", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return r.Indexes().CreateMany(ctx, indexes, opts...)
}

func (r *certificateRepo) Enables(ctx context.Context) ([]*model.Certificate, error) {
	filter := bson.D{{Key: "enabled", Value: true}}
	return r.Find(ctx, filter)
}
