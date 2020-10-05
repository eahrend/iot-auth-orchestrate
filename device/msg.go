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
	topicSplit := strings.Split(message.Topic(), "/")
	switch operation := topicSplit[len(topicSplit)-1]; operation {
	case "reverseString":
		reverseMessage(client, message)
	case "addTwo":
		addTwo(client, message)
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

func reverseMessage(client mqtt.Client, msg mqtt.Message) {
	messageData := msg.Payload()
	messageToReverseBlob := common.ReverseMessage{}
	err := json.NewDecoder(bytes.NewBuffer(messageData)).Decode(&messageToReverseBlob)
	if err != nil {
		fmt.Println(err)
		return
	}
	messageToReverse := messageToReverseBlob.MessageString
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

func addTwo(client mqtt.Client, msg mqtt.Message) {
	messageData := msg.Payload()
	addTwoMessage := common.AddTwoMessage{}
	err := json.NewDecoder(bytes.NewBuffer(messageData)).Decode(&addTwoMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	firstNumber := addTwoMessage.NumberOne
	secondNumber := addTwoMessage.NumberTwo
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
