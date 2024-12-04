package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"time"
)

var (
	Mongodburl         = "mongodb://mongoadmin:secret@127.0.0.1:27017"
	Dbname             = "mongodata"
	Dbcollectiontxdata = "txdata"
)

// mongodb
func GetMongoClient(mongodb string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongodb,
	))

	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetCollection(c *mongo.Client, dbname, collection string) *mongo.Collection {
	err := c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("error in connecting mongodb: %v", err)
	}
	database := c.Database(dbname)
	dataCollection := database.Collection(collection)
	return dataCollection
}
