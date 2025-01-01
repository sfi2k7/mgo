package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sfi2k7/blueregistryclient"
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

func etom(e bson.E) map[string]interface{} {
	m := make(map[string]interface{})
	value := e.Value

	if _, ok := value.(bson.D); ok {
		value = dtom(value.(bson.D))
	}

	m[e.Key] = value
	return m
}

func dtom(d bson.D) map[string]interface{} {
	m := make(map[string]interface{})
	for _, each := range d {
		value := each.Value

		if _, ok := value.(bson.D); ok {
			value = dtom(value.(bson.D))
		}

		if _, ok := value.(bson.E); ok {
			value = etom(value.(bson.E))
		}

		m[each.Key] = value
	}
	return m
}

func main() {
	mongoUrl := blueregistryclient.GetKeyUsingDefaultUrl("apps.migration.blueserver.blueserver")
	// if runtime.GOOS == "darwin" {
	// 	mongoUrl = "mongodb://localhost:27017"
	// }

	s, err := mgo.NewSession(mongoUrl)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	var result bson.D
	// var command = bson.D{{"find", "providers"}, {"filter", bson.E{"_id", "302420"}}}
	var command = bson.M{"hostInfo": 1}
	// cursor, err := s.DB("acoapp").RunCommandCursor(command)
	err = s.DB("acoapp").RunCommand(command, &result)
	if err != nil {
		fmt.Println("err: (run) ", err)
	}

	var m = dtom(result)

	jsoned, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println("result: ", string(jsoned))
	// fmt.Printf("m: %+v\n ", m)

	// err = cursor.All(context.Background(), &result)
	// if err != nil {
	// 	fmt.Println("err: (all)", err)
	// }

	// fmt.Printf("result: %v\n ", result)

	// var status interface{}
	// s.DB("acoapp").RunCommand(bson.M{"collStats": "providers"}, &status)
	// stats, err := s.DB("acoapp").C("providers").Stats()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("stats: %+v\n ", stats.StorageSize)
	// d := status.(bson.D)
	// m := dtom(d)
	// jsoned, _ := json.Marshal(m)

	// fmt.Println("status: ", string(jsoned))

	// jsoned, _ := json.Marshal(m)
	// fmt.Println("status: ", string(jsoned))

	// jsoned, _ := json.Marshal(status)
	// fmt.Println("status: ", string(jsoned))
	// fmt.Println("err: ", err)
	// fmt.Println("status: ", status)
	// indeces := s.DB("acoapp").C("providers").ListIndexes()
	// for _, index := range indeces {
	// 	jsoned, _ := json.Marshal(index)
	// 	fmt.Println("index:", string(jsoned))
	// }
	// return

	// RenameExample()
	// indexExample()
	// return

	// s, err := mgo.NewSession("mongodb://localhost:27017")
	// if err != nil {
	// 	panic(err)
	// }
	// defer s.Close()

	// var all []*S3Zipper
	// err = s.DB("blue").C("s3zipper").Find(bson.M{}).Select(
	// 	bson.M{
	// 		"_id":       1,
	// 		"status":    1,
	// 		"startedAt": 1,
	// 	}).Sort("-startedAt").Skip(0).Limit(10).All(&all)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("all: ", len(all))

	// for _, each := range all {
	// 	fmt.Println("each: ", each.StartedAt, each.ID)
	// }
}

type indexobject struct {
	Background       bool   `bson:"background"`
	DefaultLanguage  string `bson:"default_language"`
	Key              bson.D `bson:"key"`
	LanguageOverride string `bson:"language_override"`
	Name             string `bson:"name"`
	TextIndexVersion int    `bson:"textIndexVersion"`
	V                int    `bson:"v"`
	Weights          bson.D `bson:"weights"`
}

func RenameExample() {
	s, err := mgo.NewSession("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	err = s.DB("admin").RenameCollection("needrenamed.oldname", "needrenamed.newname")
	if err != nil {
		panic(err)
	}
}

func indexExample() {
	s, err := mgo.NewSession("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	// fmt.Println("Connected to MongoDB!", s.Ping())
	// c, err := s.DB("blueAction").C("t_keys").Indexes().List(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }
	var all = s.DB("blueAction").C("t_keys").ListIndexes()
	// err = c.All(context.TODO(), &all)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("all: ", len(all))

	for _, each := range all {
		fmt.Println("each: ", each)
	}
}
