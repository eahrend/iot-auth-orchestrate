package common

import "time"

type IotMessage struct {
	Metadata IoTMetaData `json:"metadata"`
	Message  interface{} `json:"message"`
}

// ReverseMessage takes the string and reverses the message
type ReverseMessage struct {
	MessageString string `json:"message_string"`
}

type AddTwoMessage struct {
	NumberOne int `json:"number_one"`
	NumberTwo int `json:"number_two"`
}

type IoTMetaData struct {
	Version int    `json:"version"`
	Method  string `json:"method"`
}

// IoTDeviceState contains the data about the status of
// a single device, this _should_ be a bit more strict
// but for the purposes of demonstration, I'll
// just use map[string]string
type IoTDeviceState struct {
	TimeStamp time.Time         `json:"time_stamp"`
	DeviceID  string            `json:"device_id"`
	Status    map[string]string `json:"status"`
}

type UserAuthDataKey struct {
	Users map[string]UserAuthData `json:"users"`
}

type UserAuthData struct {
	Password      string   `json:"password"`
	AuthedDevices []string `json:"authed_devices"`
}
