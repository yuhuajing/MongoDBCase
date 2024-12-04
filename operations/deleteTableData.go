package operations

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/db"
)

func DeleteDocuments(ctx context.Context, client *mongo.Client, dbs, collection string, filter bson.M) error {
	collections := db.GetCollection(client, dbs, collection)
	_, err := collections.DeleteMany(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete documents: %v", err)
	}
	return nil
}
