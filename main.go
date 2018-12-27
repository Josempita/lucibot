package main

import (
	"fmt"
	"log"
	"os"
	"time"
	 "math/rand"
	"github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
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
	min := 25.0
	max := 40.0
	r := 25.0
	for  {
		 r = min + rand.Float64() * (max - min)
		text := fmt.Sprintf("{temperature: %f}", r)
		fmt.Println(text)
		token := c.Publish("v1/devices/me/telemetry", 0, false, text)
		token.Wait()
		r = min + rand.Float64() * (max - min)
		 text = fmt.Sprintf("{humidity: %f}", r)
		token = c.Publish("v1/devices/me/telemetry", 0, false, text)
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
