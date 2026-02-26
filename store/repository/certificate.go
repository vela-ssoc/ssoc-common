package repository

import (
	"context"
	"crypto/tls"

	"github.com/vela-ssoc/ssoc-common/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Certificate interface {
	Repository[model.Certificate]
	LoadCertificate(ctx context.Context) ([]*tls.Certificate, error)
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

func (r *certificateRepo) LoadCertificate(ctx context.Context) ([]*tls.Certificate, error) {
	filter := bson.D{{Key: "enabled", Value: true}}
	dats, err := r.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	pairs := make([]*tls.Certificate, 0, len(dats))
	for _, dat := range dats {
		pair, err1 := tls.X509KeyPair([]byte(dat.PublicKey), []byte(dat.PrivateKey))
		if err1 != nil {
			return nil, err1
		}
		pairs = append(pairs, &pair)
	}

	return pairs, nil
}
