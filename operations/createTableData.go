package operations

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main/common/db"
	"main/common/table"
)

//var datas []interface{}
//var d1 = table.NFTData{
//	Status:  true,
//	Message: "BatchThree",
//	Result: table.Trans{
//		BlockNumber:      101,
//		TransactionIndex: 1,
//		Gas:              99,
//	},
//}
//datas = append(datas, d1)
//
//var d2 = table.NFTData{
//	Status:  true,
//	Message: "BatchFour",
//	Result: table.Trans{
//		BlockNumber:      102,
//		TransactionIndex: 2,
//		Gas:              97,
//	},
//}
//datas = append(datas, d2)
//err = operations.InsertMany(ctx, client, db.Dbname, db.Dbcollectiontxdata, datas)
//if err != nil {
//return
//}

func InsertDocument(ctx context.Context, client *mongo.Client, dbs, collection string, data interface{}) error {
	collections := db.GetCollection(client, dbs, collection)
	result, err := collections.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	} else {
		log.Printf("Insert data success with id = %v", result.InsertedID)
	}
	return nil
}

func InsertDocuments(ctx context.Context, client *mongo.Client, dbs, collection string, datas []interface{}) error {
	collections := db.GetCollection(client, dbs, collection)
	result, err := collections.InsertMany(ctx, datas)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	} else {
		log.Printf("Insert data success with id = %v", result.InsertedIDs)
	}
	return nil
}

func InsertOne(ctx context.Context, client *mongo.Client, dbs, collection string, data table.NFTData) error {
	collections := db.GetCollection(client, dbs, collection)
	result, err := collections.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	} else {
		log.Printf("Insert data success with id = %v", result.InsertedID)
	}
	return nil
}

func InsertMany(ctx context.Context, client *mongo.Client, dbs, collection string, datas []interface{}) error {
	collections := db.GetCollection(client, dbs, collection)
	result, err := collections.InsertMany(ctx, datas)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	} else {
		log.Printf("Insert data success with id = %v", result.InsertedIDs)
	}
	return nil
}
