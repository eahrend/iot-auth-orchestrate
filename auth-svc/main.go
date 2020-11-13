package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/eahrend/iot-auth-orchestrate/auth-svc-sql/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"os"
)

func main() {
	sqlUser := os.Getenv("MYSQL_USER")
	sqlPassword := os.Getenv("MYSQL_PASSWORD")
	sqlHost := os.Getenv("MYSQL_HOST")
	sqlDB := os.Getenv("MYSQL_DATABASE")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", sqlUser, sqlPassword, sqlHost, sqlDB))
	if err != nil {
		log.Println(fmt.Sprintf("%s:%s@%s/%s", sqlUser, sqlPassword, sqlHost, sqlDB))
		log.Fatalln(err.Error())
	}
	boil.SetDB(db)
	app := gin.Default()
	app.Use(UseSQL(db))
	app.GET("/auth", Authenticate)
	app.GET("/auth/:deviceID", AuthenticateForDeviceID)
	panic(app.Run("0.0.0.0:8081"))
}

func UseSQL(executor boil.ContextExecutor) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("boiler", executor)
		c.Next()
	}
}

type UserDevices struct {
	Name string `boil:"device_name"`
}

func Authenticate(c *gin.Context) {
	db := c.MustGet("boiler").(boil.ContextExecutor)
	username, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth {
		c.JSON(401, gin.H{
			"error": "you gotta give me a password numbnuts",
		})
		c.Abort()
		return
	}
	user, _ := models.Users(Where("username = ?", username)).One(context.Background(), db)
	if password != user.Password {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
	obj := []UserDevices{}
	err := queries.Raw("SELECT DISTINCT devices.device_name FROM user_devices JOIN users on user_devices.user_id = users.id JOIN devices on user_devices.device_id = devices.id where users.username = ?", user.Username).Bind(context.Background(), db, &obj)
	if err != nil {
		panic(err)
	}
	c.JSON(200, obj)
}

func AuthenticateForDeviceID(c *gin.Context) {
	db := c.MustGet("boiler").(boil.ContextExecutor)
	deviceID := c.Param("deviceID")
	username, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
	device, _ := models.Devices(Where("device_name = ?", deviceID)).One(context.Background(), db)
	user, _ := models.Users(Where("username = ?", username)).One(context.Background(), db)
	if password != user.Password {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
	userDeviceCheck, _ := models.UserDevices(Where("device_id = ? and user_id =?", device.ID, user.ID)).Exists(context.Background(), db)
	if !userDeviceCheck {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
	c.JSON(200, nil)
	return
}
