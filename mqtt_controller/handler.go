package mqttcontroller

import (
	"encoding/json"
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
	// res, err := http.Get("http://localhost:3000/patients")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// reBody, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("response: %s\n", reBody)

	patients, err := mh.RecordService.RetrieveRecordsByPatientID(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	patientJson, err := json.Marshal(patients)
	fmt.Println(string(patientJson))

	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	// mh.RecordService.InsertNewRecord()
}

func (mh MQTTHandler) OnConnectHandler(client MQTT.Client) {
	fmt.Println("Connected to MQTT broker")
}

func (mh MQTTHandler) ConnectionLostHandler(client MQTT.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}
