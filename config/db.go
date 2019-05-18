package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//NewSession creates a new mongo session returning a pointer to the asked database collection
func NewSession(uri, db, cl string) *mongo.Collection {

	//create a context for comunicate with mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//iniatiate the pointed client and connects to the mongo server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("uri"))
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	//pings the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	//assign a pointes collection
	col := client.Database("db").Collection("cl")

	return col
}
