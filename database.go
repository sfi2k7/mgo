package mgo

import "go.mongodb.org/mongo-driver/mongo"

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

func (d *Database) Run(cmd interface{}, result interface{}) error {
	return d.db.RunCommand(d.s.ctx, cmd).Decode(result)
}
