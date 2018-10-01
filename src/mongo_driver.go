package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var collection *mgo.Collection
var session mgo.Session

func initMongo() bool {
	//connecting with mongo db

	connString := fmt.Sprintf("%s:%s", mongoConfig.host, mongoConfig.port)
	println(connString)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{connString},
		Username: mongoConfig.username,
		Password: mongoConfig.password,
		Database: mongoConfig.database,
	})
	if err != nil {
		fmt.Println(err)
	}

	collection = session.DB(mongoConfig.database).C("sensor_data")

	return true
}

//function for save in mongoDB
func insertRecord(msg MQTTMessage) bool {
	err := collection.Insert(msg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
