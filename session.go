package mgo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	c       *mongo.Client
	ctx     context.Context
	counter int
}

func (s *Session) Close() error {
	err := s.c.Disconnect(s.ctx)
	s.c = nil
	return err
}

func (s *Session) CloseNoOp() {
	s.counter--
}

func (s *Session) SetSocketTimeout(d time.Duration) {
	fmt.Println("SetSocketTimeout is not implemented")
}

func (s *Session) DatabaseNames() ([]string, error) {
	return s.c.ListDatabaseNames(s.ctx, bson.D{})
}

func (s *Session) NumberSessionsInProgress() int {
	return s.c.NumberSessionsInProgress()
}

func (s *Session) Ping() error {
	return s.c.Ping(s.ctx, nil)
}

func (s *Session) DB(name string) *Database {
	s.counter++
	return &Database{s.c.Database(name), s}
}
