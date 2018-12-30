package sensor

type SensorType interface {
	GetValue() float64
	GetName() string
	GetRandomValue() float64
	GetMQTTValue() string
}
