package main

import (
	"os"
)

//MQTTConfig - Structure for mqtt config parameters
type MQTTConfig struct {
	host     string
	port     string
	username string
	password string
	topic    string
	qos      string
}

//MongoConfig - Structure for mongo config parameters
type MongoConfig struct {
	host     string
	port     string
	database string
	username string
	password string
}

var mongoConfig = MongoConfig{
	host:     getEnv("MONGO_HOST", "localhost"),
	port:     getEnv("MONGO_PORT", "27017"),
	database: getEnv("MONGO_DB", "airQuality"),
	username: getEnv("MONGO_USER", ""),
	password: getEnv("MONGO_PASSWORD", ""),
}

var mqttConfig = MQTTConfig{
	host:     getEnv("MQTT_HOST", "52.77.234.8"),
	port:     getEnv("MQTT_PORT", "1883"),
	username: getEnv("MQTT_USER", ""),
	password: getEnv("MQTT_PASSWORD", ""),
	topic:    getEnv("MQTT_TOPIC", "sensor_data"),
	qos:      getEnv("MQTT_QOS", "0"),
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
