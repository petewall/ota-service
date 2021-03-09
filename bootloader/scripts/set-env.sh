#!/bin/bash

echo FIRMWARE_VERSION
export FIRMWARE_VERSION=$(cat version)

echo WIFI_SSID
export WIFI_SSID=$(cat ../../secrets/pipeline-creds.json | jq -r '.["wifi-ssid"]')

echo WIFI_PASSWORD
export WIFI_PASSWORD=$(cat ../../secrets/pipeline-creds.json | jq -r '.["wifi-password"]')
