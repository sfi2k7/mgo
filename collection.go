package mgo

import (
	"context"

	"github.com/sfi2k7/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateResult = mongo.UpdateResult

func IsErrNotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}

type Collection struct {
	c *mongo.Collection
	d *Database
}

func (c *Collection) EstimatedCount() (int64, error) {
	return c.c.EstimatedDocumentCount(context.Background())
}

func (c *Collection) Find(filter interface{}) *Cursor {
	return &Cursor{cursortype: cursorTypeQuery, c: c, filter: filter}
}

func (c *Collection) FindOne(filter interface{}) *Cursor {
	return &Cursor{cursortype: cursorTypeQuery, issingle: true, c: c, filter: filter}
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

func (c *Collection) FindId(id interface{}) *Cursor {
	return &Cursor{issingle: true, c: c, filter: bson.M{"_id": id}}
}

func (c *Collection) Delete(filter interface{}) (int, error) {
	dr, err := c.c.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(dr.DeletedCount), err
}

func (c *Collection) DeleteAll(filter interface{}) (int, error) {
	dm, err := c.c.DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(dm.DeletedCount), err
}

func (c *Collection) Remove(filter interface{}) (int, error) {
	dr, err := c.c.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(dr.DeletedCount), err
}

func (c *Collection) RemoveAll(filter interface{}) (int, error) {
	dm, err := c.c.DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(dm.DeletedCount), err
}

func (c *Collection) DeleteId(id interface{}) error {
	_, err := c.c.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (c *Collection) RemoveId(id interface{}) error {
	_, err := c.c.DeleteOne(context.Background(), bson.M{"_id": id})
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

func (c *Collection) Update(filter interface{}, update interface{}) (*UpdateResult, error) {
	ur, err := c.c.UpdateMany(context.Background(), filter, update)
	return ur, err
}

func (c *Collection) UpdateOne(filter interface{}, update interface{}) (*UpdateResult, error) {
	ur, err := c.c.UpdateOne(context.Background(), filter, update)
	return ur, err
}

func (c *Collection) UpdateId(id interface{}, update interface{}) (*UpdateResult, error) {
	ur, err := c.c.UpdateByID(context.Background(), id, update)
	return ur, err
}

func (c *Collection) Upsert(filter, update interface{}) (*UpdateResult, error) {
	upsert := true
	uopt := options.UpdateOptions{Upsert: &upsert}
	ur, err := c.c.UpdateOne(context.Background(), filter, update, &uopt)
	return ur, err
}

func (c *Collection) UpsertId(id, update interface{}) (*UpdateResult, error) {
	upsert := true
	uopt := options.UpdateOptions{Upsert: &upsert}
	ur, err := c.c.UpdateByID(context.Background(), bson.M{"_id": id}, update, &uopt)
	return ur, err
}

func (c *Collection) Pipe(pipeline interface{}) *AggregateCursor {
	return &AggregateCursor{c: c.c, pipeline: pipeline}
}

func (c *Collection) Bulk() *Bulk {
	return &Bulk{c: c}
}
