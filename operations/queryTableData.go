package operations

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/common/db"
	"reflect"
)

func GetDocuments(ctx context.Context, client *mongo.Client, dbs, collection string, filter bson.M, result interface{}, opts ...*options.FindOptions) (error, []interface{}) {
	collections := db.GetCollection(client, dbs, collection)
	cur, err := collections.Find(ctx, filter, opts...)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("failed to get documents: %v", err)), nil
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {

		}
	}(cur, ctx)

	var results []interface{}
	for cur.Next(ctx) {
		// Create a new instance of the result type for each document
		res := reflect.New(reflect.TypeOf(result).Elem()).Interface()
		err := cur.Decode(res)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed to decode document: %v", err)), nil
		}
		results = append(results, res)
	}

	if err := cur.Err(); err != nil {
		return fmt.Errorf(fmt.Sprintf("cursor error: %v", err)), nil
	}

	return nil, results
}
