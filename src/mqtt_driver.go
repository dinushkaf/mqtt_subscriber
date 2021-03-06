package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//MQTTMessage - Message structure for mqtt json object
type MQTTMessage struct {
	Serial      string  `json:"SerialNo" bson:"SerialNo"`
	TimeStamp   string  `json:"Timestamp" bson:"Timestamp"`
	Temperature float32 `json:"Temp" bson:"Temp"`
	Humidity    float32 `json:"Humid" bson:"Humid"`
	PM2         float32 `json:"PM2" bson:"PM2"`
	HCHO        float32 `json:"HCHO" bson:"HCHO"`
	Ozone       float32 `json:"Ozone" bson:"Ozone"`
	CO2         float32 `json:"CO2" bson:"CO2"`
	CO          float32 `json:"CO" bson:"CO"`
	TVOC        float32 `json:"TVOC" bson:"TVOC"`
}

//MongoDocument - Message structure for mongo document object
type MongoDocument struct {
	Serial       string    `json:"SerialNo" bson:"SerialNo"`
	TimeStamp    float64   `json:"Timestamp" bson:"Timestamp"`
	Temperature  float32   `json:"Temp" bson:"Temp"`
	Humidity     float32   `json:"Humid" bson:"Humid"`
	PM2          float32   `json:"PM2" bson:"PM2"`
	HCHO         float32   `json:"HCHO" bson:"HCHO"`
	Ozone        float32   `json:"Ozone" bson:"Ozone"`
	CO2          float32   `json:"CO2" bson:"CO2"`
	CO           float32   `json:"CO" bson:"CO"`
	TVOC         float32   `json:"TVOC" bson:"TVOC"`
	ReceivedTime time.Time `json:"ReceivedTime" bson:"ReceivedTime"`
}

var mqttClient MQTT.Client
var choke = make(chan [2]string)

func initMQTT() bool {
	//set options for broker
	broker := fmt.Sprintf("tcp://%s:%s", mqttConfig.host, mqttConfig.port)
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(mqttConfig.username)
	opts.SetPassword(mqttConfig.password)
	opts.SetClientID("testClient")
	opts.SetCleanSession(false)

	//adding channel for publish handler
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	//connecting MQTT client
	mqttClient := MQTT.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(broker)
		panic(token.Error())
	}

	//subscribe to topic
	qosAsInt, errConv := strconv.Atoi(mqttConfig.qos)
	if errConv != nil {
		fmt.Println("Invalid value for QOS, using default value(0)")
		qosAsInt = 0
	}
	if token := mqttClient.Subscribe(mqttConfig.topic, byte(qosAsInt), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	return true
}

func subscribMqtt() {
	fmt.Println("Subscriber Listening...")
	for true {
		incoming := <-choke
		var msg MQTTMessage
		err := json.Unmarshal([]byte(incoming[1]), &msg)
		if err != nil {
			fmt.Printf("Error : %s", err)
		} else {
			mDoc := MQTT2Mongo(msg)
			mDoc.ReceivedTime = time.Now()
			insertRecord(mDoc)
		}
	}
}

func disconnectMQTT() {
	mqttClient.Disconnect(250)
	fmt.Println("Subscriber Disconnected")
}

//MQTT2Mongo : this will convert Mqtt msg in to mongo document
func MQTT2Mongo(msg MQTTMessage) MongoDocument {
	var tempDoc MongoDocument
	tempDoc.Serial = msg.Serial
	tempDoc.TimeStamp, _ = strconv.ParseFloat(msg.TimeStamp, 64) //time.Parse("20060102150405", msg.TimeStamp)
	tempDoc.CO2 = msg.CO2
	tempDoc.CO = msg.CO
	tempDoc.HCHO = msg.HCHO
	tempDoc.Humidity = msg.Humidity
	tempDoc.Ozone = msg.Ozone
	tempDoc.PM2 = msg.PM2
	tempDoc.Temperature = msg.Temperature
	tempDoc.TVOC = msg.TVOC
	return tempDoc
}
