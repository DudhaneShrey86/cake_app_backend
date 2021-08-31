package connection

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MainCtx context.Context
var MainClient *mongo.Client
var MainDb *mongo.Database
var CakeColl *mongo.Collection
var CategoryColl *mongo.Collection

func ConnectToDB() func() {
	uri := "mongodb://localhost:27017"
	MainCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MainClient, err := mongo.Connect(MainCtx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	closeConnection := func() {
		if err := MainClient.Disconnect(MainCtx); err != nil {
			panic(err)
		}
	}
	if err := MainClient.Ping(MainCtx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	MainDb = MainClient.Database("cake_shop")
	CakeColl = MainDb.Collection("cakes")
	CategoryColl = MainDb.Collection("categories")

	return closeConnection
}
