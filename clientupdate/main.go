package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/r3labs/sse"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	server := sse.New()
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	go publishEvents(ctx, server, client, "device-state")
	go publishEvents(ctx, server, client, "addtwo")
	go publishEvents(ctx, server, client, "reverse")
	app := gin.Default()
	app.Use(Authenticate())
	app.Handle(http.MethodGet, "/events", gin.WrapF(server.HTTPHandler))
	panic(app.Run("0.0.0.0:8082"))
}

func publishEvents(ctx context.Context, server *sse.Server, client *pubsub.Client, subscriptionID string) {
	sub := client.Subscription(subscriptionID)
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		server.Publish(m.Attributes["deviceId"], &sse.Event{
			Data: m.Data,
		})
	})
	if err != nil {
		panic(err)
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
