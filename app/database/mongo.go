package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

var DB *mongo.Database

var MongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority",
	os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"),
	os.Getenv("MONGO_PORT"))

func InitMongo() {
	fmt.Println("MONGO URL: ", MongoURI)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoURI))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("[MONGO CONNECTION] Successfully connected to host [%s] and database [%s]", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_DB")))
	DB = client.Database(os.Getenv("MONGO_DB"))
}
