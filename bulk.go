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

type BulkResult struct {
	MatchedCount  int64
	ModifiedCount int64
	UpsertedCount int64
	UpsertedIDs   map[int64]interface{}
	DeletedCount  int64
	InsertedCount int64
}

func (b *Bulk) Run() (*BulkResult, error) {
	br, err := b.c.c.BulkWrite(context.Background(), b.models)
	if err != nil {
		return nil, err
	}

	return &BulkResult{
		MatchedCount:  br.MatchedCount,
		ModifiedCount: br.ModifiedCount,
		UpsertedCount: br.UpsertedCount,
		UpsertedIDs:   br.UpsertedIDs,
		DeletedCount:  br.DeletedCount,
		InsertedCount: br.InsertedCount,
	}, nil
}
