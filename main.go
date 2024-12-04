package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"main/common/db"
	"main/common/table"
	"main/operations"
)

func main() {
	client, err := db.GetMongoClient(db.Mongodburl)
	if err != nil {
		log.Fatalf("err in conn MonggoDB: %v", err)
	}

	var ctx = context.Background()
	var datas []interface{}
	var d1 = table.NFTData{
		Status:  true,
		Message: "BatchThree",
		Result: table.Trans{
			BlockNumber:      101,
			TransactionIndex: 1,
			Gas:              99,
		},
	}
	datas = append(datas, d1)

	var d2 = table.NFTData{
		Status:  true,
		Message: "BatchFour",
		Result: table.Trans{
			BlockNumber:      102,
			TransactionIndex: 2,
			Gas:              97,
		},
	}
	datas = append(datas, d2)
	err = operations.InsertMany(ctx, client, db.Dbname, db.Dbcollectiontxdata, datas)
	if err != nil {
		return
	}
	//var filter = bson.M{"blocknumber": bson.D{{"$lt", 100}}}
}
