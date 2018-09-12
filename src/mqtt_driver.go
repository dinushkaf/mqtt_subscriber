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
	Serial       string  `json:"serial"`
	TimeStamp    string  `json:"timestamp"`
	Temperature  float32 `json:"temp"`
	Humidity     float32 `json:"humid"`
	PM2          float32 `json:"pm2"`
	HCHCHO       float32 `json:"hchcho"`
	Ozone        float32 `json:"ozone"`
	CO2          float32 `json:"co2"`
	TVOC         float32 `json:"tvoc"`
	ReceivedTime string  `json:"receivedTime"`
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
			fmt.Println(msg)
			msg.ReceivedTime = time.Now().Format("20060102150405")
			insertRecord(msg)
		}
	}
}

func disconnectMQTT() {
	mqttClient.Disconnect(250)
	fmt.Println("Subscriber Disconnected")
}
