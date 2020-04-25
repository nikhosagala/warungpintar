package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"net/url"
)

var mqttHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func mqttConnect(clientId string, uri *url.URL) MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetClientID(clientId)
	opts.SetDefaultPublishHandler(mqttHandler)
	newClient := MQTT.NewClient(opts)
	if token := newClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return newClient
}

func mqttClient(clientId string) MQTT.Client {
	uri, _ := url.Parse("http://localhost:1883")
	client := mqttConnect(clientId, uri)
	return client
}

func listen(topic string) {
	client := mqttClient("sub")
	client.Subscribe(topic, 0, mqttHandler)
}
