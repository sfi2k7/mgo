package mgo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Cursor struct {
	c *mongo.Cursor
}

func (c *Cursor) All(result interface{}) error {
	if c.c == nil {
		fmt.Println("cursor is nil")
	}

	return c.c.All(context.Background(), result)
}
