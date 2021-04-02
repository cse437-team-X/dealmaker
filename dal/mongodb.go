package dal

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Mongo *mongo.Client

var ItemCollection *mongo.Collection

func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	Mongo = client
	db := Mongo.Database("dealmaker")
	//CredUserCollection = db.Collection("cred_user")
	ItemCollection = db.Collection("items")
}
