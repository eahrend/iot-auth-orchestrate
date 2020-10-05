#!/usr/bin/env bash

for deviceID in device1 device2 device3
do
  openssl ecparam -genkey -name prime256v1 -noout -out ${PWD}/deviceKeys/${deviceID}/ec_private.pem
  openssl ec -in ec_private.pem -pubout -out ${PWD}/deviceKeys/${deviceID}/ec_public.pem
done


