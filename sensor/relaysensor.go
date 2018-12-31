package sensor

import (
	"fmt"
	"math/rand"
)

type RelaySensor struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	UseRandom bool    `json:"useRandom"`
	State     bool    `json:"state"`
}

func (l RelaySensor) GetName() string {
	return l.Name
}

func (l RelaySensor) GetMQTTValue() string {
	text := fmt.Sprintf("{%s: %t}", l.Name, l.State)
	return text
}

func (l RelaySensor) GetValue() float64 {
	return 0.0
}

func (l RelaySensor) GetRandomValue() float64 {
	min := 25.0
	max := 40.0
	r := min + rand.Float64()*(max-min)
	return r
}

func (l RelaySensor) GetState() bool {
	return l.State
}
