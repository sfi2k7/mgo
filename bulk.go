package mgo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type Bulk struct {
	c      *Collection
	models []mongo.WriteModel
}

var ErrInvalidBulkUpdatePairs = errors.New("invalid bulk update pairs")

func (b *Bulk) Update(pairs ...interface{}) error {

	if len(pairs)%2 != 0 {
		return ErrInvalidBulkUpdatePairs
	}

	for i := 0; i < len(pairs); i += 2 {
		pair := pairs[i : i+2]
		b.models = append(b.models, mongo.NewUpdateOneModel().SetFilter(pair[0]).SetUpdate(pair[1]))
	}

	return nil
}

func (b *Bulk) Upsert(pairs ...interface{}) error {

	if len(pairs)%2 != 0 {
		return ErrInvalidBulkUpdatePairs
	}

	for i := 0; i < len(pairs); i += 2 {
		pair := pairs[i : i+2]
		b.models = append(b.models, mongo.NewUpdateOneModel().SetFilter(pair[0]).SetUpdate(pair[1]).SetUpsert(true))
	}

	return nil
}

func (b *Bulk) Insert(docs ...interface{}) error {

	for _, doc := range docs {
		b.models = append(b.models, mongo.NewInsertOneModel().SetDocument(doc))
	}

	return nil
}

func (b *Bulk) Delete(filters ...interface{}) error {

	for _, filter := range filters {
		b.models = append(b.models, mongo.NewDeleteOneModel().SetFilter(filter))
	}

	return nil
}

type BulkResult struct {
	MatchedCount  int
	ModifiedCount int
	UpsertedCount int
	UpsertedIDs   map[int64]interface{}
	DeletedCount  int
	InsertedCount int
}

func (b *Bulk) Run() (*BulkResult, error) {
	br, err := b.c.c.BulkWrite(context.Background(), b.models)
	if err != nil {
		return nil, err
	}

	return &BulkResult{
		MatchedCount:  int(br.MatchedCount),
		ModifiedCount: int(br.ModifiedCount),
		UpsertedCount: int(br.UpsertedCount),
		UpsertedIDs:   br.UpsertedIDs,
		DeletedCount:  int(br.DeletedCount),
		InsertedCount: int(br.InsertedCount),
	}, nil
}
