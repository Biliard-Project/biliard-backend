package mqttcontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/Biliard-Project/biliard-backend/models"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	Broker   = "tcp://test.mosquitto.org:1883"
	Topic    = "/Capstone/msg"
	ClientID = "biliard"
)

type MQTTHandler struct {
	RecordService      *models.RecordService
	PatientScanService *models.PatientScanService
}

type SensorOutput struct {
	R         uint8   `json:"r"`
	G         uint8   `json:"g"`
	B         uint8   `json:"b"`
	C         uint8   `json:"c"`
	HeartRate float64 `json:"hr"`
	Oxygen    float64 `json:"o"`
}

type MLInput struct {
	Red   uint8 `json:"Red"`
	Green uint8 `json:"Green"`
	Blue  uint8 `json:"Blue"`
}

type MLOutput struct {
	Prediction float64 `json:"prediction"`
}

func (mh MQTTHandler) MessagePubHandler(client MQTT.Client, msg MQTT.Message) {
	var sensorOutput SensorOutput
	fmt.Println(string(msg.Payload()))

	err := json.Unmarshal(msg.Payload(), &sensorOutput)
	if err != nil {
		fmt.Println("jsoning 0")
		fmt.Printf("Error Unmarshalling %s\n", msg.Payload())
		return
	}

	mlInput := MLInput{
		Red:   sensorOutput.R,
		Green: sensorOutput.G,
		Blue:  sensorOutput.B,
	}

	jsonData, err := json.Marshal(mlInput)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:5005/predict", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("3")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("4")
		fmt.Println(err)
		return
	}

	var mlOutput MLOutput
	err = json.Unmarshal(body, &mlOutput)
	if err != nil {
		fmt.Println("5")
		fmt.Println(err)
		return
	}
	mlOutput.Prediction = math.Round(mlOutput.Prediction*100) / 100

	patient, err := mh.PatientScanService.Get()
	if err != nil {
		fmt.Println("6")
		fmt.Println(err)
		return
	}

	insertedpatient, err := mh.RecordService.InsertNewRecord(patient.ID, models.JSONTime(time.Now()), mlOutput.Prediction, sensorOutput.Oxygen, sensorOutput.HeartRate)
	if err != nil {
		fmt.Println("7")
		fmt.Println(err)
		return
	}
	fmt.Println(insertedpatient)
}

func (mh MQTTHandler) OnConnectHandler(client MQTT.Client) {
	fmt.Println("Connected to MQTT broker")
}

func (mh MQTTHandler) ConnectionLostHandler(client MQTT.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}
