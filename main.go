package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Josempita/lucibot/sensor"
	"github.com/eclipse/paho.mqtt.golang"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func getSensorValue(s sensor.TemperatureSensor) {

}

type RpcMessage struct {
	Method string `json:"method"`
	Params bool   `json:"params"`
}

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetUsername("Z7jEJIXHKAH0hippTBuH")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("v1/devices/me/telemetry", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println("Error!")
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//subscribe to RPC calls from the server

	var callback MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		topic := msg.Topic()
		payload := msg.Payload()
		s := string(payload[:])
		messageID := msg.MessageID()
		request := "v1/devices/me/rpc/request/"
		requestID := topic[len(request):len(topic)]

		var payloadMessage RpcMessage
		if err := json.Unmarshal(payload, &payloadMessage); err != nil {
			panic(err)
		}

		fmt.Printf("TOPIC: %s\n", topic)
		fmt.Printf("MSG: %s\n", payload)
		fmt.Printf("payload: %s\n", s)
		fmt.Printf("MSG ID: %d\n", messageID)
		fmt.Printf("MSG Method: %s\n", payloadMessage.Method)
		fmt.Printf("MSG Method: %s\n", requestID)
		response, err := json.Marshal(payloadMessage)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client.Publish("v1/devices/me/rpc/response/"+requestID, 0, false, string(response))

	}

	if token2 := c.Subscribe("v1/devices/me/rpc/request/+", 0, callback); token2.Wait() && token2.Error() != nil {
		fmt.Println("Error!")
		fmt.Println(token2.Error())
		os.Exit(1)
	}

	tempSensor := sensor.TemperatureSensor{Name: "temperature", Value: 0.0, UseRandom: true, State: true}
	humidSensor := sensor.HumiditySensor{Name: "humidity", Value: 0.0, UseRandom: true, State: true}
	relay := sensor.RelaySensor{Name: "relay", Value: 0.0, UseRandom: false, State: true}

	for {

		token := c.Publish("v1/devices/me/telemetry", 0, false, tempSensor.GetMQTTValue())
		token.Wait()

		token = c.Publish("v1/devices/me/telemetry", 0, false, humidSensor.GetMQTTValue())
		token.Wait()

		token = c.Publish("v1/devices/me/telemetry", 0, false, relay.GetMQTTValue())
		token.Wait()

		time.Sleep(1 * time.Second)

	}

	time.Sleep(6 * time.Second)

	if token := c.Unsubscribe("v1/devices/me/telemetry"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}
