package operations

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/db"
)

func UpdateDocument(ctx context.Context, client *mongo.Client, dbs, collection string, filter bson.M, update bson.M) error {
	fmt.Println("up")
	collections := db.GetCollection(client, dbs, collection)
	_, err := collections.UpdateMany(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %v", err)
	}
	return nil
}
