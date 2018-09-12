package main

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var mongoClient *mongo.Client
var err error
var database *mongo.Database
var collection *mongo.Collection

func initMongo() bool {
	//setting connection string
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoConfig.username, mongoConfig.password,
		mongoConfig.host, mongoConfig.port)
	if mongoConfig.username == "" && mongoConfig.password == "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s",
			mongoConfig.host, mongoConfig.port)
	} else if mongoConfig.username == "" || mongoConfig.password == "" {
		fmt.Println("Please provide MONGO_USER and MONGO_PASSWORD")
		return false
	}
	fmt.Println("Connection String : " + connectionString)

	//connecting with mongo db
	mongoClient, err = mongo.Connect(context.Background(), connectionString)
	if err != nil {
		fmt.Println("Mongo connection error occured!")
		fmt.Printf("Connection String : %s\n", connectionString)
		return false
	}

	//setting database and collection
	database = mongoClient.Database(mongoConfig.database)
	collection = database.Collection("sensor_data")

	return true
}

//function for save in mongoDB
func insertRecord(msg MQTTMessage) bool {
	result, err := collection.InsertOne(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("Inserted : %s", result.InsertedID)
	return true
}
