package mgo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type AggregateCursor struct {
	cu       *mongo.Cursor
	c        *mongo.Collection
	pipeline interface{}
}

func (c *AggregateCursor) All(result interface{}) error {
	cursor, err := c.c.Aggregate(context.Background(), c.pipeline)
	if err != nil {
		return err
	}

	err = cursor.All(context.TODO(), result)
	c.cu = nil
	c.c = nil
	return err
}

func (c *AggregateCursor) Run() (*mongo.Cursor, error) {
	cursor, err := c.c.Aggregate(context.Background(), c.pipeline)
	c.cu = cursor
	return cursor, err
}

func (c *AggregateCursor) Close() error {
	if c.cu == nil {
		return errors.New("cursor is nil")
	}

	c.c = nil
	c.cu = nil
	return c.cu.Close(context.TODO())
}

func (c *AggregateCursor) Next(result interface{}) bool {
	if c.cu == nil {
		return false
	}

	return c.cu.Next(context.TODO())
}

func (c *AggregateCursor) Decode(result interface{}) error {
	if c.cu == nil {
		return errors.New("cursor is nil")
	}

	return c.cu.Decode(result)
}

func (c *AggregateCursor) Err() error {
	if c.cu == nil {
		return errors.New("cursor is nil")
	}

	return c.cu.Err()
}

func (c *AggregateCursor) ID() int64 {
	if c.cu == nil {
		return 0
	}

	return c.cu.ID()
}

func (c *AggregateCursor) RemainingBatchLength() int {
	if c.cu == nil {
		return 0
	}

	return c.cu.RemainingBatchLength()
}
