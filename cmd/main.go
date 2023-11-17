package main

import (
	"fmt"
	"time"

	"github.com/sfi2k7/mgo"
	"github.com/sfi2k7/mgo/bson"
)

/*{
    "_id" : "f396d438-117d-49ea-a1ce-75118ee5ec14",
    "currentFileCount" : 0,
    "currentSourceFile" : "",
    "error" : "open /var/s3zipper/f396d438-117d-49ea-a1ce-75118ee5ec14.zip: no such file or directory",
    "publicFilePath" : "",
    "startedAt" : ISODate("2023-08-11T12:51:29.649Z"),
    "status" : "failed",
    "totalFileCount" : 2,
    "totalProcessTook" : "10.934125ms",
    "zipFileSize" : NumberLong(0),
    "reason" : "Could not create local Zipfile"
}*/

type S3Zipper struct {
	ID                string    `bson:"_id"`
	CurrentFileCount  int       `bson:"currentFileCount"`
	CurrentSourceFile string    `bson:"currentSourceFile"`
	Error             string    `bson:"error"`
	PublicFilePath    string    `bson:"publicFilePath"`
	StartedAt         time.Time `bson:"startedAt"`
	Status            string    `bson:"status"`
	TotalFileCount    int       `bson:"totalFileCount"`
	TotalProcessTook  string    `bson:"totalProcessTook"`
	ZipFileSize       int64     `bson:"zipFileSize"`
	Reason            string    `bson:"reason"`
}

func main() {
	s, err := mgo.NewSession("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	var all []*S3Zipper
	err = s.DB("blue").C("s3zipper").Find(bson.M{}).Select(bson.M{"_id": 1, "status": 1, "startedAt": 1}).Sort("-startedAt").Skip(0).Limit(10).All(&all)
	if err != nil {
		panic(err)
	}

	fmt.Println("all: ", len(all))

	for _, each := range all {
		fmt.Println("each: ", each.StartedAt, each.ID)
	}
}
