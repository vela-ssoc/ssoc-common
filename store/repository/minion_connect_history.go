package repository

import (
	"context"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MinionConnectHistory interface {
	Repository[model.MinionConnectHistory]
}

func NewMinionConnectHistory(db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) MinionConnectHistory {
	repo := NewRepository[model.MinionConnectHistory](db, opts...)

	return &minionConnectHistoryRepo{Repository: repo}
}

type minionConnectHistoryRepo struct {
	Repository[model.MinionConnectHistory]
}

func (r *minionConnectHistoryRepo) CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "minion_id", Value: 1}}},
		{Keys: bson.D{{Key: "tunnel_stat.connected_at", Value: -1}}},
	}

	return r.Indexes().CreateMany(ctx, indexes, opts...)
}
