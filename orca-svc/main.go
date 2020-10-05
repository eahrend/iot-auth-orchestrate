package main

import (
	iot "cloud.google.com/go/iot/apiv1"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"
	"io/ioutil"
	"net/http"
)

func main() {
	ctx := context.Background()
	deviceManager, err := iot.NewDeviceManagerClient(ctx)
	if err != nil {
		panic(err)
	}
	app := gin.Default()
	app.Use(UseCloudIoT(deviceManager))
	app.Use(Authenticate())
	app.POST("/add/:deviceID", AddTwo)
	app.POST("/reverse/:deviceID", ReverseString)
	panic(app.Run("0.0.0.0:8080"))
}

func sendCommand(c *gin.Context, subfolder string) {
	deviceID, _ := c.Params.Get("deviceID")
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	ciot := c.MustGet("iot").(*iot.DeviceManagerClient)
	req := &iotpb.SendCommandToDeviceRequest{
		Name:       deviceID,
		BinaryData: bodyBytes,
		Subfolder:  subfolder,
	}
	_, err := ciot.SendCommandToDevice(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, nil)
	return
}

func AddTwo(c *gin.Context) {
	sendCommand(c, "addTwo")
	return
}

func ReverseString(c *gin.Context) {
	sendCommand(c, "reverseString")
	return
}

func UseCloudIoT(svc *iot.DeviceManagerClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("deviceManager", svc)
		c.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceID := c.Param("deviceID")
		user, password, _ := c.Request.BasicAuth()
		client := &http.Client{}
		authURL := "0.0.0.0:8081"
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/auth/%s", authURL, deviceID), nil)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		req.SetBasicAuth(user, password)
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		c.JSON(resp.StatusCode, nil)
		c.Next()
	}
}
