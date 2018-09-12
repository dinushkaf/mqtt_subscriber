FROM golang:1.9

MAINTAINER Dinushka Fernando (dinimz@live.com)

# install dependencies
RUN go get github.com/eclipse/paho.mqtt.golang
RUN go get github.com/mongodb/mongo-go-driver/mongo

# env
ENV MONGO_HOST 192.169.0.1

# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/mqttSubscriber src/*.go

ENTRYPOINT ["/app/initialize.sh"]
