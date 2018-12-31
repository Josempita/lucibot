package sensor

import (
	"fmt"
	"math/rand"
)

type TemperatureSensor struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	UseRandom bool    `json:"useRandom"`
	State     bool    `json:"state"`
}

func (t TemperatureSensor) GetName() string {
	return t.Name
}

func (t TemperatureSensor) GetMQTTValue() string {
	mqttValue := t.GetValue()
	if t.UseRandom {
		mqttValue = t.GetRandomValue()
	}
	text := fmt.Sprintf("{%s: %f}", t.Name, mqttValue)
	return text
}

func (t TemperatureSensor) GetValue() float64 {
	return t.Value
}

func (t TemperatureSensor) GetRandomValue() float64 {
	min := 25.0
	max := 40.0
	r := min + rand.Float64()*(max-min)
	return r
}

func (t TemperatureSensor) GetState() bool {
	return t.State
}
