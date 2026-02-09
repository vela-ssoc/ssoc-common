package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Broker interface {
	Repository[model.Broker]
	FindBySecret(ctx context.Context, secret string, opts ...options.Lister[options.FindOneOptions]) (*model.Broker, error)
}

func NewBroker(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) Broker {
	repo := NewRepository[model.Broker](db, opts...)

	return &brokerRepo{Repository: repo}
}

type brokerRepo struct {
	Repository[model.Broker]
}

func (r *brokerRepo) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "name", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "secret", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	return r.Indexes().CreateMany(ctx, indexes, opts...)
}

func (r *brokerRepo) FindBySecret(ctx context.Context, secret string, opts ...options.Lister[options.FindOneOptions]) (*model.Broker, error) {
	return r.FindOne(ctx, bson.D{{Key: "secret", Value: secret}}, opts...)
}
