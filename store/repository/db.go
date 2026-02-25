package repository

import (
	"context"
	"log/slog"
	"reflect"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database interface {
	Database() *mongo.Database
	CreateIndex(ctx context.Context) error

	Broker() Broker
	BrokerConnectHistory() BrokerConnectHistory
	LokiConfig() LokiConfig
	PyroscopeConfig() PyroscopeConfig
	VictoriaMetricsConfig() VictoriaMetricsConfig
}

func NewDB(db *mongo.Database, log *slog.Logger) Database {
	return &mongoDB{
		db:                    db,
		log:                   log,
		broker:                NewBroker(db),
		brokerConnectHistory:  NewBrokerConnectHistory(db),
		lokiConfig:            NewLokiConfig(db),
		pyroscopeConfig:       NewPyroscopeConfig(db),
		victoriaMetricsConfig: NewVictoriaMetricsConfig(db),
	}
}

type mongoDB struct {
	db  *mongo.Database
	log *slog.Logger

	broker                Broker
	brokerConnectHistory  BrokerConnectHistory
	lokiConfig            LokiConfig
	pyroscopeConfig       PyroscopeConfig
	victoriaMetricsConfig VictoriaMetricsConfig
}

func (r *mongoDB) Database() *mongo.Database                    { return r.db }
func (r *mongoDB) Broker() Broker                               { return r.broker }
func (r *mongoDB) BrokerConnectHistory() BrokerConnectHistory   { return r.brokerConnectHistory }
func (r *mongoDB) LokiConfig() LokiConfig                       { return r.lokiConfig }
func (r *mongoDB) PyroscopeConfig() PyroscopeConfig             { return r.pyroscopeConfig }
func (r *mongoDB) VictoriaMetricsConfig() VictoriaMetricsConfig { return r.victoriaMetricsConfig }

func (r *mongoDB) CreateIndex(ctx context.Context) error {
	rv := reflect.ValueOf(r)
	for i := range rv.NumMethod() {
		mv := rv.Method(i)
		name, ic := r.reflectCall(mv, mv.Type())
		if ic == nil {
			continue
		}

		attrs := []any{"collection_name", name}
		indexes, err := ic.CreateIndex(ctx)
		if err != nil {
			attrs = append(attrs, "error", err)
			r.log.Error("索引创建错误", attrs...)
			return err
		}

		if len(indexes) == 0 {
			r.log.Debug("无需创建索引", attrs...)
		} else {
			attrs = append(attrs, "indexes", indexes)
			r.log.Info("索引创建完毕", attrs...)
		}
	}

	return nil
}

func (r *mongoDB) reflectCall(mv reflect.Value, mt reflect.Type) (string, indexCreator) {
	if mt.NumIn() != 0 || mt.NumOut() != 1 {
		return "", nil
	}

	rets := mv.Call([]reflect.Value{})
	if len(rets) != 1 {
		return "", nil
	}
	ret := rets[0]
	if ret.IsNil() || !ret.IsValid() {
		return "", nil
	}

	val := ret.Interface()
	ic, ok := val.(indexCreator)
	if !ok {
		return "", nil
	}
	coll := ret.Type().Name()
	if ni, yes := val.(interface{ Name() string }); yes {
		coll = ni.Name()
	}

	return coll, ic
}

type indexCreator interface {
	CreateIndex(context.Context, ...options.Lister[options.CreateIndexesOptions]) ([]string, error)
}
