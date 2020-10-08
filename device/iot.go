package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// createJWT creates a Cloud IoT Core JWT for the given project id.
// algorithm can be one of ["RS256", "ES256"].
func createJWT(projectID string, privateKeyPath string, algorithm string, expiration time.Duration) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  projectID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(expiration).Unix(),
	}

	keyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(algorithm), claims)

	switch algorithm {
	case "RS256":
		privKey, _ := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
		return token.SignedString(privKey)
	case "ES256":
		privKey, err := jwt.ParseECPrivateKeyFromPEM(keyBytes)
		if err != nil {
			log.Fatalln("Failed to parse private key from pem: %v", err)
		}
		return token.SignedString(privKey)
	}

	return "", errors.New("Cannot find JWT algorithm. Specify 'ES256' or 'RS256'")
}

func connectToIoT(projectID string, region string, registryID string, deviceID string, privateKeyPath string, algorithm string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	const (
		mqttBrokerURL   = "tls://mqtt.googleapis.com:8883"
		protocolVersion = 4 // corresponds to MQTT 3.1.1
	)
	// ROOT_PEM is the location of this file
	// https://cloud.google.com/iot/docs/how-tos/mqtt-bridge#downloading_mqtt_server_certificates
	/*
		rootPEM,_ := ioutil.ReadFile(os.Getenv("ROOT_PEM"))
		roots := x509.NewCertPool()
		ok := roots.AppendCertsFromPEM([]byte(rootPEM))
		if !ok {
			panic("failed to parse root certificate")
		}
	*/
	jwt, err := createJWT(projectID, privateKeyPath, algorithm, 60)
	if err != nil {
		log.Fatalf("failed to create jwt: %v", err)
	}
	clientID := fmt.Sprintf("projects/%s/locations/%s/registries/%s/devices/%s", projectID, region, registryID, deviceID)
	log.Println("ClientID:", clientID)
	opts := mqtt.NewClientOptions()
	// if your device doesn't have these certs, this is how
	// you add them to the TLS
	//opts.SetTLSConfig(&tls.Config{RootCAs: roots})
	opts.SetTLSConfig(&tls.Config{})
	opts.AddBroker(mqttBrokerURL)
	opts.SetClientID(clientID)
	opts.SetUsername("unused")
	opts.SetPassword(jwt)
	opts.SetProtocolVersion(protocolVersion)
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(fmt.Sprintf("/devices/%s/commands/#", deviceID), 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := mqtt.NewClient(opts)
	log.Println("Creating client connection")
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		//panic(token.Error())
		log.Fatalln("Token error on connection:", token.Error())
	} else {
		log.Printf("Connected!")
	}
	// this just simulates the device doing device stuff
	// without the user constantly asking for an update
	// i.e. push notification
	go gocronRunner(client)
	<-c
}

func gocronRunner(client mqtt.Client) {
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(1).Minutes().Do(updateStateWithRandomData, client)
	s1.StartAsync()
}
