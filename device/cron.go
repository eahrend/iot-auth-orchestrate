package main

import (
	"fmt"
	"github.com/eahrend/iot-auth-orchestrate/common"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"os"
	"time"
)

func updateStateWithRandomData(client mqtt.Client) {
	status := common.IoTDeviceState{
		TimeStamp: time.Now(),
		DeviceID:  os.Getenv("DEVICE_ID"),
		Status: map[string]string{
			"uuid": uuid.New().String(),
		},
	}
	stateTopic := fmt.Sprintf("/devices/%s/state", os.Getenv("DEVICE_ID"))
	client.Publish(stateTopic, 1, false, status)
}
