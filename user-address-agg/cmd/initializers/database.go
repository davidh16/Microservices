package initializers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Client *mongo.Client

func ConnectToMongo() {
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://user:password123456789@mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer Client.Disconnect(ctx)
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
