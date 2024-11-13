package mgo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	cursorTypeQuery       = 1
	cursorTypeAggregation = 2
)

type Cursor struct {
	c          *Collection
	issingle   bool
	filter     interface{}
	limit      int
	skip       int
	sort       interface{}
	projection interface{}
	cursortype int
}

func (c *Cursor) Limit(limit int) *Cursor {
	c.limit = limit
	return c
}

func (c *Cursor) Skip(skip int) *Cursor {
	c.skip = skip
	return c
}

func (c *Cursor) Count() (int, error) {
	if c.c == nil {
		return 0, errors.New("cursor is closed")
	}

	n, err := c.c.c.CountDocuments(context.Background(), c.filter)

	return int(n), err
}

func (c *Cursor) EstimatedCount() (int, error) {
	if c.c == nil {
		return 0, errors.New("cursor is closed")
	}

	n, err := c.c.c.EstimatedDocumentCount(context.Background())

	return int(n), err
}

func (c *Cursor) Sort(sort ...string) *Cursor {

	if len(sort) == 0 {
		return c
	}

	var d bson.D
	for _, each := range sort {
		if each == "" {
			continue
		}

		if each[0] == '-' {
			d = append(d, bson.E{Key: each[1:], Value: -1})
			continue
		}

		d = append(d, bson.E{Key: each, Value: 1})
	}

	c.sort = d
	return c
}

func (c *Cursor) Select(projection interface{}) *Cursor {
	c.projection = projection
	return c
}

func (c *Cursor) One(result interface{}) error {
	if c.c == nil {
		return errors.New("cursor is closed")
	}

	var opts []*options.FindOneOptions
	if c.projection != nil {
		opts = append(opts, options.FindOne().SetProjection(c.projection))
	}

	if c.sort != nil {
		opts = append(opts, options.FindOne().SetSort(c.sort))
	}

	if c.skip > 0 {
		opts = append(opts, options.FindOne().SetSkip(int64(c.skip)))
	}

	err := c.c.c.FindOne(context.Background(), c.filter, opts...).Decode(result)

	c.c = nil
	return err
}

func (c *Cursor) All(result interface{}) error {
	if c.c == nil {
		return errors.New("cursor is closed")
	}

	var opts []*options.FindOptions

	if c.limit > 0 {
		opts = append(opts, options.Find().SetLimit(int64(c.limit)))
	}

	if c.skip > 0 {
		opts = append(opts, options.Find().SetSkip(int64(c.skip)))
	}

	if c.sort != nil {
		opts = append(opts, options.Find().SetSort(c.sort))
	}

	if c.projection != nil {
		opts = append(opts, options.Find().SetProjection(c.projection))
	}

	cur, err := c.c.c.Find(context.Background(), c.filter, opts...)

	if err != nil {
		return err
	}

	err = cur.All(context.Background(), result)
	c.c = nil
	return err
}
