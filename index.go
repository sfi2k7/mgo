package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type indexobject struct {
	Background       bool   `bson:"background,omitempty"`
	DefaultLanguage  string `bson:"default_language,omitempty"`
	Key              bson.D `bson:"key,omitempty"`
	LanguageOverride string `bson:"language_override,omitempty"`
	Name             string `bson:"name,omitempty"`
	TextIndexVersion int32  `bson:"textIndexVersion,omitempty"`
	V                int    `bson:"v,omitempty"`
	Weights          bson.D `bson:"weights,omitempty"`
	Unique           bool   `bson:"unique,omitempty"`
}

func (c *Collection) ListIndexes() []*indexobject {
	list, err := c.c.Indexes().List(context.TODO())
	if err != nil {
		return nil
	}

	var all []*indexobject
	err = list.All(context.TODO(), &all)
	if err != nil {
		return nil
	}

	return all
}

func (c *Collection) CreateIndexSimple(name string, keys interface{}) error {
	_, err := c.c.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetName(name),
	})
	return err
}

func (c *Collection) CreateIndex(index indexobject) error {
	indexOptions := options.Index()

	if index.Background {
		indexOptions.SetBackground(true)
	}

	if index.DefaultLanguage != "" {
		indexOptions.SetDefaultLanguage(index.DefaultLanguage)
	}

	if index.LanguageOverride != "" {
		indexOptions.SetLanguageOverride(index.LanguageOverride)
	}

	if index.TextIndexVersion != 0 {
		indexOptions.SetTextVersion(index.TextIndexVersion)
	}

	if index.Weights != nil {
		indexOptions.SetWeights(index.Weights)
	}

	if index.Name != "" {
		indexOptions.SetName(index.Name)
	}

	if index.Unique {
		indexOptions.SetUnique(true)
	}

	_, err := c.c.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    index.Key,
		Options: indexOptions,
	})
	return err
}
