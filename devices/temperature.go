package devices

import (
	core "github.com/PhilGruber/dimmy/core"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type Temperature struct {
	Device
}

func MakeTemperature(config core.DeviceConfig) Temperature {
	t := Temperature{}
	t.MqttTopic = config.Topic

	t.Current = 0
	t.Type = "temperature"
	return t
}

func (t *Temperature) GetMin() int {
	return 0
}

func (t *Temperature) GetMax() int {
	return 99
}

func NewTemperature(config core.DeviceConfig) *Temperature {
	s := MakeTemperature(config)
	return &s
}

func (t *Temperature) PublishValue(mqtt mqtt.Client) {
}

func (t *Temperature) ProcessRequest(request core.SwitchRequest) {
}

func (t *Temperature) UpdateValue() (float64, bool) {
	return 0, false
}

func (t *Temperature) GetMqttStateTopic() string {
	return t.MqttTopic
}

func (t *Temperature) getMessageHandler(channel chan core.SwitchRequest, temperature DeviceInterface) mqtt.MessageHandler {
	log.Println("Subscribing to " + temperature.GetMqttTopic())
	return func(client mqtt.Client, mqttMessage mqtt.Message) {
		payload := string(mqttMessage.Payload())
		log.Println("Received new temperature: " + string(payload))

		if request, ok := temperature.GenerateRequest(payload[:]); ok {
			channel <- request
		}
	}
}
