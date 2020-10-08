# iot-auth-orchestrate
IoT Authorization and Orchestration Example

# Project Layout

## iot-auth-orchestrate/service-account.json
Ignored by git, service account key that has access to the following roles in your project:

- roles/cloudiot.viewer
- roles/cloudiot.deviceController

## auth-svc
This http service acts as the auth layer between orca-svc and the IoT-API, right now it's just basic auth, but can be extended to use JWT or whatever

## common
Models/structs/functions shared between the other repos.

## device
The IoT "device" as a container, every 5 minutes it sends a random uuid to the device state topic, as well as receiving commands and updating specific topics

## devicekeys
Contains the public/private keys of devices. Ignored by git.

docker-compose mounts the private key a volume

Folder structure is as follows:
```
iot-auth-orchestrate/
|
|---devicekeys/
|   |
|   |--device1/
|   |   |--public key
|   |   |--private key
|   |
|   |--device2/
|   |   |--public key
|   |   |--private key
        
``` 

## docs
Documentation, images, charts, etc.

## iot-gcf
Consumers for the device events, could be a GCF, but it will probably be a container

## orca-svc
Gateway to the IoT API for users to interact with.

## web
the last thing I'm going to do cause I hate doing frontend work

## terraform
If you wanna stand this up in GCP, this should stand up all you need. Note that there are billable components.  

