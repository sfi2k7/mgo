package mgo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	db *mongo.Database
	s  *Session
}

func (d *Database) C(name string) *Collection {
	return &Collection{d.db.Collection(name), d}
}

func (d *Database) DropDatabase() error {
	return d.db.Drop(d.s.ctx)
}

func (d *Database) Name() string {
	return d.db.Name()
}

func (d *Database) RunCommand(cmd, result interface{}) error {
	return d.db.RunCommand(d.s.ctx, cmd).Decode(result)
}

func (d *Database) RunCommandCursor(cmd interface{}) (*mongo.Cursor, error) {
	return d.db.RunCommandCursor(d.s.ctx, cmd)
}

func (d *Database) Session() *Session {
	return d.s
}

func (d *Database) RenameCollection(oldName, newName string) error {
	singleResult := d.db.RunCommand(
		context.TODO(),
		bson.D{{Key: "renameCollection", Value: fmt.Sprintf("%s.%s", d.db.Name(), oldName)},
			{Key: "to", Value: fmt.Sprintf("%s.%s", d.db.Name(), newName)},
		})

	return singleResult.Err()
}

func (d *Database) CollectionNames() ([]string, error) {
	return d.db.ListCollectionNames(d.s.ctx, bson.D{})
}

func (d *Database) Run(cmd interface{}, result interface{}) error {
	return d.db.RunCommand(d.s.ctx, cmd).Decode(result)
}
