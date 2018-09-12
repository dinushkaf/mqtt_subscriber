package main

import (
	"os"
)

func main() {
	// Init Mondo Connection
	if !initMongo() {
		os.Exit(0)
	}

	// Init MQTT Connection
	if !initMQTT() {
		os.Exit(0)
	}

	// Start subscriber loop
	subscribMqtt()
}
