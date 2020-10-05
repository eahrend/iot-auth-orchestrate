package main

import (
	"encoding/json"
	"fmt"
	"github.com/eahrend/iot-auth-orchestrate/common"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

func updateState(client mqtt.Client, status common.IoTDeviceState) {
	stateTopic := fmt.Sprintf("/devices/%s/state", os.Getenv("DEVICE_ID"))
	b, _ := json.Marshal(&status)
	if token := client.Publish(stateTopic, 1, false, string(b)); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
}
