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
type RegEx = primitive.Regex
type Encoder = bson.Encoder
type Decoder = bson.Decoder
type RawValue = bson.RawValue
type RawElement = bson.RawElement

func ObjectIdHex() string {
	return primitive.NewObjectID().Hex()
}

func IsValidObjectIdHex(s string) bool {
	_, err := primitive.ObjectIDFromHex(s)
	return err == nil
}

func Marshal(val interface{}) ([]byte, error) {
	return bson.Marshal(val)
}

func Unmarshal(data []byte, val interface{}) error {
	return bson.Unmarshal(data, val)
}

func MarshalExtJSON(val interface{}, canonical, escapeHTML bool) ([]byte, error) {
	return bson.MarshalExtJSON(val, canonical, escapeHTML)
}

func UnmarshalExtJSON(data []byte, canonical bool, val interface{}) error {
	return bson.UnmarshalExtJSON(data, canonical, val)
}
