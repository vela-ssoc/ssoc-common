package repository

import (
	"context"
	"iter"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CollectionNamer interface {
	// CollectionName 集合名。
	// 注意：该方法不应该存在复杂的逻辑，不应该依赖 struct 内部的状态。
	CollectionName() string
}

type Pages[T any] struct {
	Page    int64 `json:"page"`
	Size    int64 `json:"size"`
	Count   int64 `json:"count"`
	Records []*T  `json:"records"`
}

func EmptyPages[T any](size int64) *Pages[T] {
	return &Pages[T]{
		Page:    1,
		Size:    size,
		Records: []*T{},
	}
}

type Repository[T any] interface {
	Name() string
	Database() *mongo.Database
	Collection() *mongo.Collection
	BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...options.Lister[options.BulkWriteOptions]) (*mongo.BulkWriteResult, error)
	InsertOne(ctx context.Context, doc *T, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)
	InsertMany(ctx context.Context, docs []*T, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error)
	DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter any, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error)
	UpdateByID(ctx context.Context, id, update any, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)
	UpdateOne(ctx context.Context, filter, update any, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter, update any, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error)
	ReplaceOne(ctx context.Context, filter, replacement any, opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error)
	Aggregate(ctx context.Context, pipe any, opts ...options.Lister[options.AggregateOptions]) ([]*T, error)
	CountDocuments(ctx context.Context, filter any, opts ...options.Lister[options.CountOptions]) (int64, error)
	EstimatedDocumentCount(ctx context.Context, opts ...options.Lister[options.EstimatedDocumentCountOptions]) (int64, error)
	Distinct(ctx context.Context, fieldName string, filter any, opts ...options.Lister[options.DistinctOptions]) *mongo.DistinctResult
	Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) ([]*T, error)
	FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) (*T, error)
	FindOneAndDelete(ctx context.Context, filter any, opts ...options.Lister[options.FindOneAndDeleteOptions]) (*T, error)
	FindOneAndReplace(ctx context.Context, filter, replacement any, opts ...options.Lister[options.FindOneAndReplaceOptions]) (*T, error)
	FindOneAndUpdate(ctx context.Context, filter, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) (*T, error)
	Watch(ctx context.Context, pipeline any, opts ...options.Lister[options.ChangeStreamOptions]) (*mongo.ChangeStream, error)
	Indexes() mongo.IndexView
	SearchIndexes() mongo.SearchIndexView
	Drop(ctx context.Context, opts ...options.Lister[options.DropCollectionOptions]) error
	FindByID(ctx context.Context, id any, opts ...options.Lister[options.FindOneOptions]) (*T, error)
	DeleteByID(ctx context.Context, id any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)
	DistinctString(ctx context.Context, fieldName string, filter any, opts ...options.Lister[options.DistinctOptions]) ([]string, error)
	DistinctObjectID(ctx context.Context, fieldName string, filter any, opts ...options.Lister[options.DistinctOptions]) ([]bson.ObjectID, error)
	AggregateTo(ctx context.Context, pipe, result any, opts ...options.Lister[options.AggregateOptions]) error
	Page(ctx context.Context, filter any, page, size int64, opts ...options.Lister[options.FindOptions]) (*Pages[T], error)
	All(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) iter.Seq2[*T, error]
	CreateIndex(ctx context.Context, opts ...options.Lister[options.CreateIndexesOptions]) ([]string, error)
}

func NewRepository[T CollectionNamer](db *mongo.Database, opts ...options.Lister[options.CollectionOptions]) Repository[T] {
	var t T
	name := t.CollectionName()

	return NewBaseRepository[T](db, name, opts...)
}

func NewBaseRepository[T any](db *mongo.Database, collName string, opts ...options.Lister[options.CollectionOptions]) Repository[T] {
	coll := db.Collection(collName, opts...)

	return &baseRepository[T]{coll: coll}
}

type baseRepository[T any] struct {
	coll *mongo.Collection
}

func (r *baseRepository[T]) Name() string                  { return r.coll.Name() }
func (r *baseRepository[T]) Database() *mongo.Database     { return r.coll.Database() }
func (r *baseRepository[T]) Collection() *mongo.Collection { return r.coll }

func (r *baseRepository[T]) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...options.Lister[options.BulkWriteOptions]) (*mongo.BulkWriteResult, error) {
	return r.coll.BulkWrite(ctx, models, opts...)
}

func (r *baseRepository[T]) InsertOne(ctx context.Context, doc *T, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	return r.coll.InsertOne(ctx, doc, opts...)
}

func (r *baseRepository[T]) InsertMany(ctx context.Context, docs []*T, opts ...options.Lister[options.InsertManyOptions]) (*mongo.InsertManyResult, error) {
	return r.coll.InsertMany(ctx, docs, opts...)
}

func (r *baseRepository[T]) DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	return r.coll.DeleteOne(ctx, filter, opts...)
}

func (r *baseRepository[T]) DeleteMany(ctx context.Context, filter any, opts ...options.Lister[options.DeleteManyOptions]) (*mongo.DeleteResult, error) {
	return r.coll.DeleteMany(ctx, filter, opts...)
}

func (r *baseRepository[T]) UpdateByID(ctx context.Context, id, update any, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	return r.coll.UpdateByID(ctx, id, update, opts...)
}

func (r *baseRepository[T]) UpdateOne(ctx context.Context, filter, update any, opts ...options.Lister[options.UpdateOneOptions]) (*mongo.UpdateResult, error) {
	return r.coll.UpdateOne(ctx, filter, update, opts...)
}

func (r *baseRepository[T]) UpdateMany(ctx context.Context, filter, update any, opts ...options.Lister[options.UpdateManyOptions]) (*mongo.UpdateResult, error) {
	return r.coll.UpdateMany(ctx, filter, update, opts...)
}

func (r *baseRepository[T]) ReplaceOne(ctx context.Context, filter, replacement any, opts ...options.Lister[options.ReplaceOptions]) (*mongo.UpdateResult, error) {
	return r.coll.ReplaceOne(ctx, filter, replacement, opts...)
}

func (r *baseRepository[T]) Aggregate(ctx context.Context, pipe any, opts ...options.Lister[options.AggregateOptions]) ([]*T, error) {
	cur, err := r.coll.Aggregate(ctx, pipe, opts...)
	if err != nil {
		return nil, err
	}

	return r.cursorAll(ctx, cur)
}

func (r *baseRepository[T]) AggregateTo(ctx context.Context, pipe, result any, opts ...options.Lister[options.AggregateOptions]) error {
	cur, err := r.coll.Aggregate(ctx, pipe, opts...)
	if err != nil {
		return err
	}

	return r.cursorTo(ctx, cur, result)
}

func (r *baseRepository[T]) CountDocuments(ctx context.Context, filter any, opts ...options.Lister[options.CountOptions]) (int64, error) {
	return r.coll.CountDocuments(ctx, filter, opts...)
}

func (r *baseRepository[T]) EstimatedDocumentCount(ctx context.Context, opts ...options.Lister[options.EstimatedDocumentCountOptions]) (int64, error) {
	return r.coll.EstimatedDocumentCount(ctx, opts...)
}

func (r *baseRepository[T]) Distinct(ctx context.Context, fieldName string, filter any, opts ...options.Lister[options.DistinctOptions]) *mongo.DistinctResult {
	return r.coll.Distinct(ctx, fieldName, filter, opts...)
}

func (r *baseRepository[T]) Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) ([]*T, error) {
	cur, err := r.coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return r.cursorAll(ctx, cur)
}

func (r *baseRepository[T]) FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) (*T, error) {
	return r.decodeOne(r.coll.FindOne(ctx, filter, opts...))
}

func (r *baseRepository[T]) FindOneAndDelete(ctx context.Context, filter any, opts ...options.Lister[options.FindOneAndDeleteOptions]) (*T, error) {
	return r.decodeOne(r.coll.FindOneAndDelete(ctx, filter, opts...))
}

func (r *baseRepository[T]) FindOneAndReplace(ctx context.Context, filter, replacement any, opts ...options.Lister[options.FindOneAndReplaceOptions]) (*T, error) {
	return r.decodeOne(r.coll.FindOneAndReplace(ctx, filter, replacement, opts...))
}

func (r *baseRepository[T]) FindOneAndUpdate(ctx context.Context, filter, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) (*T, error) {
	return r.decodeOne(r.coll.FindOneAndUpdate(ctx, filter, update, opts...))
}

func (r *baseRepository[T]) Watch(ctx context.Context, pipeline any, opts ...options.Lister[options.ChangeStreamOptions]) (*mongo.ChangeStream, error) {
	return r.coll.Watch(ctx, pipeline, opts...)
}

func (r *baseRepository[T]) Indexes() mongo.IndexView             { return r.coll.Indexes() }
func (r *baseRepository[T]) SearchIndexes() mongo.SearchIndexView { return r.coll.SearchIndexes() }

func (r *baseRepository[T]) Drop(ctx context.Context, opts ...options.Lister[options.DropCollectionOptions]) error {
	return r.coll.Drop(ctx, opts...)
}

func (r *baseRepository[T]) FindByID(ctx context.Context, id any, opts ...options.Lister[options.FindOneOptions]) (*T, error) {
	return r.FindOne(ctx, bson.E{Key: "_id", Value: id}, opts...)
}

func (r *baseRepository[T]) DeleteByID(ctx context.Context, id any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	return r.DeleteOne(ctx, bson.E{Key: "_id", Value: id}, opts...)
}

func (r *baseRepository[T]) DistinctString(ctx context.Context, fieldName string, filter any, opts ...options.Lister[options.DistinctOptions]) ([]string, error) {
	var ss []string
	if err := r.Distinct(ctx, fieldName, filter, opts...).Decode(&ss); err != nil {
		return nil, err
	}

	return ss, nil
}

func (r *baseRepository[T]) DistinctObjectID(ctx context.Context, fieldName string, filter any, opts ...options.Lister[options.DistinctOptions]) ([]bson.ObjectID, error) {
	var ss []bson.ObjectID
	if err := r.Distinct(ctx, fieldName, filter, opts...).Decode(&ss); err != nil {
		return nil, err
	}

	return ss, nil
}

func (r *baseRepository[T]) Page(ctx context.Context, filter any, page, size int64, opts ...options.Lister[options.FindOptions]) (*Pages[T], error) {
	page, size = r.clampPageSize(page, size)
	count, err := r.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	} else if count == 0 {
		return EmptyPages[T](size), nil
	}

	lastPage, skip, limit := r.paginate(page, size, count)
	opt := options.Find().SetSkip(skip).SetLimit(limit)
	opts = append(opts, opt)
	cur, err := r.coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	records, err := r.cursorAll(ctx, cur)
	if err != nil {
		return nil, err
	}
	ret := &Pages[T]{
		Page:    lastPage,
		Size:    size,
		Count:   count,
		Records: records,
	}

	return ret, nil
}

func (r *baseRepository[T]) All(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) iter.Seq2[*T, error] {
	return func(yield func(*T, error) bool) {
		cur, err := r.coll.Find(ctx, filter, opts...)
		if err != nil {
			yield(nil, err)
			return
		}
		//goland:noinspection GoUnhandledErrorResult
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			t := new(T)
			if err = cur.Decode(t); err != nil {
				yield(nil, err)
				return
			}
			if !yield(t, nil) {
				return
			}
		}

		if err = cur.Err(); err != nil {
			yield(nil, err)
		}
	}
}

// CreateIndex 桩代码，由各个实际 repo 实现。
func (r *baseRepository[T]) CreateIndex(context.Context, ...options.Lister[options.CreateIndexesOptions]) ([]string, error) {
	return nil, nil
}

func (r *baseRepository[T]) decodeOne(sr *mongo.SingleResult) (*T, error) {
	t := new(T)
	if err := sr.Decode(t); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *baseRepository[T]) cursorAll(ctx context.Context, cur *mongo.Cursor) ([]*T, error) {
	var ts []*T
	if err := r.cursorTo(ctx, cur, &ts); err != nil {
		return nil, err
	}

	return ts, nil
}

func (*baseRepository[T]) cursorTo(ctx context.Context, cur *mongo.Cursor, result any) error {
	//goland:noinspection GoUnhandledErrorResult
	defer cur.Close(ctx)

	return cur.All(ctx, result)
}

// paginate 分页计算，如果页码过大导致超出页数，则会保留最后一页的内容，并返回修正后的最后页码。
func (r *baseRepository[T]) paginate(page, size, count int64) (fixedPage, skip, limit int64) {
	if maximum := (count + size - 1) / size; maximum > 0 && maximum < page {
		page = maximum
	}
	skip = (page - 1) * size

	return page, size, skip
}

// clampPageSize 对输入的 page size 参数做区间限制处理。
func (r *baseRepository[T]) clampPageSize(page, size int64) (int64, int64) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	return page, size
}
