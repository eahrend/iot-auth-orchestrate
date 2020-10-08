package main

import (
	"encoding/json"
	"fmt"
	"github.com/eahrend/iot-auth-orchestrate/common"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func updateStateWithRandomData(client mqtt.Client) {
	log.Println("Updating State with Random Data")
	status := common.IoTDeviceState{
		TimeStamp: time.Now(),
		DeviceID:  os.Getenv("DEVICE_ID"),
		Status: map[string]string{
			"uuid": uuid.New().String(),
		},
	}
	b, _ := json.Marshal(&status)
	stateTopic := fmt.Sprintf("/devices/%s/state", os.Getenv("DEVICE_ID"))
	token := client.Publish(stateTopic, 1, false, string(b))
	if token.Error() != nil {
		log.Fatalln("token error:", token.Error())
	} else {
		log.Println("random data updated")
	}
}
