package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewSession(u string) (*Session, error) {
	ctx := context.Background()
	var err error
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(u))
	if err != nil {
		return nil, err
	}

	return &Session{c, ctx}, nil
}

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer func() {
// 		if err = client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	fmt.Println("Connected to MongoDB!", client.Ping(ctx, nil))

// 	cursor, err := client.Database("blue").Collection("s3zipper").Find(ctx, bson.M{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	var all []bson.M
// 	err = cursor.All(ctx, &all)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("All documents: ", len(all))

// }
