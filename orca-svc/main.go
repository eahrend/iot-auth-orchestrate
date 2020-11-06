package main

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"

	_ "golang.org/x/oauth2/google"
	"google.golang.org/api/cloudiot/v1"
)

func main() {
	ctx := context.Background()
	client, err := cloudiot.NewService(ctx)
	if err != nil {
		panic(err)
	}
	app := gin.Default()
	config := map[string]string{
		"projectID":  os.Getenv("PROJECT_ID"),
		"region":     os.Getenv("REGION"),
		"registryID": os.Getenv("REGISTRY_ID"),
	}
	app.Use(useConfig(config))
	app.Use(UseCloudIoT(client))
	app.Use(Authenticate())
	app.POST("/devices/:deviceID/commands", genericCommand)
	app.GET("/devices", GetDevices)
	app.POST("/add/:deviceID", AddTwo)
	app.POST("/reverse/:deviceID", ReverseString)
	panic(app.Run("0.0.0.0:8080"))
}

func useConfig(config map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}

func genericCommand(c *gin.Context) {
	sendCommand(c, "")
}

func GetDevices(c *gin.Context) {
	deviceID, _ := c.Params.Get("deviceID")
	ciot := c.MustGet("iot").(*cloudiot.Service)
	req := ciot.Projects.Locations.Registries.Devices.Get(deviceID)
	resp, _ := req.Do()
	b, _ := resp.State.MarshalJSON()
	c.JSON(200, string(b))
}

func sendCommand(c *gin.Context, subfolder string) {
	deviceID, _ := c.Params.Get("deviceID")
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	ciot := c.MustGet("iot").(*cloudiot.Service)
	config := c.MustGet("config").(map[string]string)
	req := cloudiot.SendCommandToDeviceRequest{
		BinaryData: b64.StdEncoding.EncodeToString(bodyBytes),
		Subfolder:  subfolder,
	}
	name := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", config["projectID"], config["region"], config["registryID"], deviceID)
	_, err := ciot.Projects.Locations.Registries.Devices.SendCommandToDevice(name, &req).Do()
	if err != nil {
		log.Println("Error from sending command to device:", err.Error())
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, map[string]string{
		"status": "sent",
	})
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

func UseCloudIoT(svc *cloudiot.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("iot", svc)
		c.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authSvc := os.Getenv("AUTH_URL")
		deviceID := c.Param("deviceID")
		user, password, _ := c.Request.BasicAuth()
		client := &http.Client{}
		authURL := fmt.Sprintf("http://%s:8081", authSvc)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/auth/%s", authURL, deviceID), nil)
		if err != nil {
			log.Println("error from auth request", err.Error())
			c.JSON(500, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		req.SetBasicAuth(user, password)
		_, err = client.Do(req)
		if err != nil {
			log.Println("error from auth client do", err.Error())
			c.JSON(500, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
