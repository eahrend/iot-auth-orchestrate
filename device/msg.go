package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eahrend/iot-auth-orchestrate/common"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"strconv"
	"strings"
	"time"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	ciot := common.IotMessage{}
	buf := bytes.NewBuffer(message.Payload())
	json.NewDecoder(buf).Decode(&ciot)
	switch operation := ciot.Metadata.Method; operation {
	case "ReverseMessage":
		reverseMessage(client, ciot)
	case "AddTwo":
		addTwo(client, ciot)
	default:
		fmt.Println("Method not supported")
	}
}
func Reverse(orig string) string {
	var c []string = strings.Split(orig, "")

	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}

	return strings.Join(c, "")
}

func reverseMessage(client mqtt.Client, msg common.IotMessage) {
	messageToReverse := msg.Message.(common.ReverseMessage).MessageString
	status := common.IoTDeviceState{
		TimeStamp: time.Now(),
		DeviceID:  os.Getenv("DEVICE_ID"),
		Status: map[string]string{
			"executingFunc":    "reverseMessage",
			"messageToReverse": messageToReverse,
			"status":           "processing",
		},
	}
	updateState(client, status)
	topicName := fmt.Sprintf("/devices/%s/events/reverse", os.Getenv("DEVICE_ID"))
	m := Reverse(messageToReverse)
	client.Publish(topicName, 1, false, m)
	status.Status = map[string]string{
		"executingFunc":    "reverseMessage",
		"messageToReverse": messageToReverse,
		"status":           "complete",
	}
	updateState(client, status)
}

func addTwo(client mqtt.Client, msg common.IotMessage) {
	firstNumber := msg.Message.(common.AddTwoMessage).NumberOne
	secondNumber := msg.Message.(common.AddTwoMessage).NumberTwo
	status := common.IoTDeviceState{
		TimeStamp: time.Now(),
		DeviceID:  os.Getenv("DEVICE_ID"),
		Status: map[string]string{
			"executingFunc": "addTwo",
			"firstNumber":   strconv.Itoa(firstNumber),
			"secondNumber":  strconv.Itoa(secondNumber),
			"status":        "processing",
		},
	}
	updateState(client, status)
	sum := firstNumber + secondNumber
	topicName := fmt.Sprintf("/devices/%s/events/addTwo", os.Getenv("DEVICE_ID"))
	client.Publish(topicName, 1, false, sum)
	status = common.IoTDeviceState{
		TimeStamp: time.Now(),
		DeviceID:  os.Getenv("DEVICE_ID"),
		Status: map[string]string{
			"executingFunc": "addTwo",
			"firstNumber":   strconv.Itoa(firstNumber),
			"secondNumber":  strconv.Itoa(secondNumber),
			"status":        "complete",
		},
	}
	updateState(client, status)
}
