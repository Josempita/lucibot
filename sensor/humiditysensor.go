package sensor

import (
	"fmt"
	"math/rand"
)

type HumiditySensor struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	UseRandom bool    `json:"useRandom"`
}

func (h HumiditySensor) GetName() string {
	return h.Name
}

func (t HumiditySensor) GetMQTTValue() string {
	mqttValue := t.GetValue()
	if t.UseRandom {
		mqttValue = t.GetRandomValue()
	}
	text := fmt.Sprintf("{%s: %f}", t.Name, mqttValue)
	return text
}

func (h HumiditySensor) GetValue() float64 {
	return h.Value
}

func (h HumiditySensor) GetRandomValue() float64 {
	min := 25.0
	max := 40.0
	r := min + rand.Float64()*(max-min)
	return r
}
