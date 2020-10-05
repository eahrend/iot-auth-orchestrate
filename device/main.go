package main

import "os"

func main() {
	projectID := os.Getenv("PROJECT_ID")
	region := os.Getenv("REGION")
	registryID := os.Getenv("REGISTRY_ID")
	deviceID := os.Getenv("DEVICE_ID")
	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
	algorithm := os.Getenv("ALGORITHM")
	connectToIoT(projectID, region, registryID, deviceID, privateKeyPath, algorithm)
}
