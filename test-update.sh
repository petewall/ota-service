#!/usr/bin/env bash

FIRMWARE=$1
if [[ "${FIRMWARE}" == "" ]] ; then
  FIRMWARE=test
fi
VERSION=$2
if [[ "${VERSION}" == "" ]] ; then
  VERSION=0.0.1
fi
MAC_ADDRESS=$3
if [[ "${MAC_ADDRESS}" == "" ]] ; then
  MAC_ADDRESS=my-test-device-mac
fi

set -x
curl --verbose --header "x-esp8266-sta-mac: ${MAC_ADDRESS}" "http://localhost:8266/api/update?firmware=${FIRMWARE}&version=${VERSION}"
