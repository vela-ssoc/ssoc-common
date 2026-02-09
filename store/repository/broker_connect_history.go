package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BrokerConnectHistory interface {
	Repository[model.BrokerConnectHistory]
}

func NewBrokerConnectHistory(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) BrokerConnectHistory {
	repo := NewRepository[model.BrokerConnectHistory](db, opts...)

	return &brokerConnectHistoryRepo{Repository: repo}
}

type brokerConnectHistoryRepo struct {
	Repository[model.BrokerConnectHistory]
}

func (r *brokerConnectHistoryRepo) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "broker_id", Value: 1}}},
		{Keys: bson.D{{Key: "tunnel_stat.connected_at", Value: -1}}},
	}

	return r.Indexes().CreateMany(ctx, indexes, opts...)
}
