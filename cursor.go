package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Cursor struct {
	c          *Collection
	filter     interface{}
	limit      int64
	skip       int64
	sort       interface{}
	projection interface{}
}

func (c *Cursor) Limit(limit int64) *Cursor {
	c.limit = limit
	return c
}

func (c *Cursor) Skip(skip int64) *Cursor {
	c.skip = skip
	return c
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

func (c *Cursor) All(result interface{}) error {
	var opts []*options.FindOptions

	if c.limit > 0 {
		opts = append(opts, options.Find().SetLimit(c.limit))
	}

	if c.skip > 0 {
		opts = append(opts, options.Find().SetSkip(c.skip))
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
	return cur.All(context.Background(), result)
}
