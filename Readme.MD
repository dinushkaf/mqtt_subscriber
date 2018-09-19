#Introduction
This will dockerize MQTT Subscriber - Data logger written in GO Lang.
All the sources are in the *src* folder.
*Dockerfile* contains the instructions needed to dockerize.
*initialize.sh* file is entrypoint.
This image will compile go code and build it in the build folder and run it in a endless loop.

#Requirements
* Docker-CE
* GO Lang(if you need to test code before dockerize)
* MQTT broker that has been deployed somewhere(or locally)
* MongoDB that has been deployed somewhere(or locally)

#Parameters
Change the following environmental variables according to your need.(see deploy section to see how to change them)
1. Parameters Related to MongoDB
 * **MONGO_HOST** 	Default : localhost 	- ip of the MongoDB Connection
 * **MONGO_PORT** 	Default : 27017 	- port of the MongoDB Connection
 * **MONGO_DB**		Default : airQuality 	- name of the Database
 * **MONGO_USER**	Default : 		- set username
 * **MONGO_PASSWORD**	Default :		- set password
2. Parameters Realted to MQTT
 * **MQTT_HOST** 	Default : 52.77.234.8 	- ip of the MQTT Broker
 * **MQTT_PORT** 	Default : 1883 		- port of the MQTT Broker
 * **MQTT_USER** 	Default :  		- set username
 * **MQTT_PASSWORD** 	Default :  		- set passsword
 * **MQTT_TOPIC** 	Default : sensor_data 	- MQTT topic to subscribe
 * **MQTT_QOS** 	Default : 0 		- MQTT Quality of service(leave this to its default)

#Deploy
* Open a terminal window at master folder
* dockerize the image by `docker build -t mqtt-logger .`
* you can run the container with `sudo docker run -d mqtt-logger`
 * use -it switches if you need to run the image iteractively
 * set environment variables with -e switch. Eg: `sudo docker run -e VAR1=1254 -e VAR2="MQTT Logger" -d mqtt-logger`
* use `docker ps -a` to see if it is up and running
* now publish a data to the broker with the given topic name and check if it saves in the mongodb
