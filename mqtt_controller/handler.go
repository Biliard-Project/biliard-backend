package mqttcontroller

import (
	"fmt"

	"github.com/Biliard-Project/biliard-backend/models"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	Broker   = "tcp://test.mosquitto.org:1883"
	Topic    = "/Capstone/msg"
	ClientID = "biliard"
)

type MQTTHandler struct {
	RecordService *models.RecordService
}

func (mh MQTTHandler) MessagePubHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	// mh.RecordService.InsertNewRecord()
}

func (mh MQTTHandler) OnConnectHandler(client MQTT.Client) {
	fmt.Println("Connected to MQTT broker")
}

func (mh MQTTHandler) ConnectionLostHandler(client MQTT.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}
