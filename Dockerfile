FROM golang:1.9

MAINTAINER Dinushka Fernando (dinimz@live.com)

# install dependencies
RUN go get github.com/eclipse/paho.mqtt.golang
RUN go get gopkg.in/mgo.v2

# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/mqttSubscriber src/*.go

ENTRYPOINT ["/app/initialize.sh"]
