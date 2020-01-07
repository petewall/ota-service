#!/usr/bin/env bash

FIRMWARE=$1
if [[ "${FIRMWARE}" == "" ]] ; then
  FIRMWARE=test
fi
VERSION=$2
if [[ "${VERSION}" == "" ]] ; then
  VERSION=0.0.1
fi

set -x
mkdir -p ./data/firmware/${FIRMWARE}/${VERSION}
touch ./data/firmware/${FIRMWARE}/${VERSION}/${FIRMWARE}_${VERSION}.bin
