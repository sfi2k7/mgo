package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	c   *mongo.Client
	ctx context.Context
}

func (s *Session) Close() error {
	return s.c.Disconnect(s.ctx)
}

func (s *Session) DatabaseNames() ([]string, error) {
	return s.c.ListDatabaseNames(s.ctx, nil)
}

func (s *Session) NumberSessionsInProgress() int {
	return s.c.NumberSessionsInProgress()
}

func (s *Session) Ping() error {
	return s.c.Ping(s.ctx, nil)
}

func (s *Session) DB(name string) *Database {
	return &Database{s.c.Database(name), s}
}
