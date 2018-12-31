package sensor

type SensorType interface {
	GetValue() float64
	GetState() bool
	GetName() string
	GetRandomValue() float64
	GetMQTTValue() string
}
