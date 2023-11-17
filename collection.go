package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	c *mongo.Collection
	d *Database
}

func (c *Collection) Find(filter interface{}) *Cursor {
	return &Cursor{c: c, filter: filter}
}

func (c *Collection) FindOne(filter interface{}, result interface{}) error {
	single := c.c.FindOne(context.Background(), filter)
	return single.Decode(result)
}

func (c *Collection) InsertOne(doc interface{}) error {
	_, err := c.c.InsertOne(context.Background(), doc)
	return err
}

func (c *Collection) Insert(docs ...interface{}) error {
	_, err := c.c.InsertMany(context.Background(), docs)
	return err
}

func (c *Collection) Count(filter interface{}) (int64, error) {
	return c.c.CountDocuments(context.Background(), filter)
}

func (c *Collection) DeleteOne(filter interface{}) error {
	_, err := c.c.DeleteOne(context.Background(), filter)
	return err
}

func (c *Collection) Delete(filter interface{}) error {
	_, err := c.c.DeleteMany(context.Background(), filter)
	return err
}

func (c *Collection) Distinct(key string, filter interface{}) ([]interface{}, error) {
	return c.c.Distinct(context.Background(), key, filter)
}

func (c *Collection) DropCollection() error {
	return c.c.Drop(context.Background())
}

func (c *Collection) EstimatedDocumentCount() (int64, error) {
	return c.c.EstimatedDocumentCount(context.Background())
}

func (c *Collection) Update(filter interface{}, update interface{}) error {
	_, err := c.c.UpdateMany(context.Background(), filter, update)
	return err
}

func (c *Collection) UpdateOne(filter interface{}, update interface{}) error {
	_, err := c.c.UpdateOne(context.Background(), filter, update)
	return err
}

func (c *Collection) Bulk() *Bulk {
	return &Bulk{c: c}
}
