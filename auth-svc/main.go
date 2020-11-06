package main

import (
	"bytes"
	"encoding/json"
	"github.com/eahrend/iot-auth-orchestrate/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	app := gin.Default()
	app.Use(UseRedis(rdb))
	app.GET("/auth/:deviceID", Authenticate)
	//app.GET("/auth/me",GetDevices)
	panic(app.Run("0.0.0.0:8081"))
}

/*
// Gets data about the user
func GetDevices(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
	rdb := c.MustGet("redis").(*redis.Client)
	userInfoBytes, err := rdb.Get("users").Bytes()
	if err != nil {
		log.Println("getting user bytes:", err.Error())
		c.JSON(500, gin.H{
			"error": err,
		})
		c.Abort()
		return
	}
	userInfo := common.UserAuthDataKey{}
	err = json.NewDecoder(bytes.NewBuffer(userInfoBytes)).Decode(&userInfo)
	if err != nil {
		log.Println("json decoding user information", err.Error())
		c.JSON(500, gin.H{
			"error": err,
		})
		c.Abort()
		return
	}
	if val, ok := userInfo.Users[user]; ok {
		if val.Password != password {
			c.JSON(401, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
	} else {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
	c.JSON(200,userInfo.Users[userID].AuthedDevices)
}

*/

func UseRedis(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Set("redis", rdb)
		c.Next()
	}
}

func Authenticate(c *gin.Context) {
	deviceID := c.Param("deviceID")
	user, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
	rdb := c.MustGet("redis").(*redis.Client)
	userInfoBytes, err := rdb.Get("users").Bytes()
	if err != nil {
		log.Println("getting user bytes:", err.Error())
		c.JSON(500, gin.H{
			"error": err,
		})
		c.Abort()
		return
	}
	userInfo := common.UserAuthDataKey{}
	err = json.NewDecoder(bytes.NewBuffer(userInfoBytes)).Decode(&userInfo)
	if err != nil {
		log.Println("json decoding user information", err.Error())
		c.JSON(500, gin.H{
			"error": err,
		})
		c.Abort()
		return
	}
	if val, ok := userInfo.Users[user]; ok {
		if val.Password != password {
			c.JSON(401, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
		deviceCheck := false
		for _, device := range val.AuthedDevices {
			if device == deviceID {
				deviceCheck = true
				break
			}
		}
		if deviceCheck {
			c.JSON(200, gin.H{
				"status": "authorized",
			})
			return
		} else {
			c.JSON(401, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
	} else {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
}
