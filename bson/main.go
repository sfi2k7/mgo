package bson

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type M = bson.M
type D = bson.D
type E = bson.E
type A = bson.A
type Raw = bson.Raw

func ObjectIdHex() string {
	return primitive.NewObjectID().Hex()
}
