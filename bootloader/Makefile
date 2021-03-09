.PHONY: build clean test

clean:
	rm -rf .pio

SRC_FILES := $(shell find . -path ./test -prune -false -o -name "*.cpp" -o -name "*.h")
TEST_FILES := $(shell find ./test -name "*.cpp" -o -name "*.h")

test: build ${TEST_FILES}
	pio test --environment ota_bootloader --verbose

FIRMWARE_VERSION=$(shell cat version)
WIFI_SSID=$(shell cat ../../secrets/pipeline-creds.json | jq -r '.["wifi-ssid"]')
WIFI_PASSWORD=$(shell cat ../../secrets/pipeline-creds.json | jq -r '.["wifi-password"]')

.pio/build/ota_bootloader/firmware.bin: platformio.ini ${SRC_FILES}
	FIRMWARE_VERSION="${FIRMWARE_VERSION}" \
	OTA_HOST="http://wallserver.local:8266" \
	WIFI_SSID="${WIFI_SSID}" \
	WIFI_PASSWORD="${WIFI_PASSWORD}" \
		pio run --environment ota_bootloader

build: .pio/build/ota_bootloader/firmware.bin

check:
	FIRMWARE_VERSION="${FIRMWARE_VERSION}" \
	OTA_HOST="http://wallserver.local:8266" \
	WIFI_SSID="${WIFI_SSID}" \
	WIFI_PASSWORD="${WIFI_PASSWORD}" \
		pio check --environment ota_bootloader --verbose --skip-packages

upload:
	FIRMWARE_VERSION="${FIRMWARE_VERSION}" \
	OTA_HOST="http://wallserver.local:8266" \
	WIFI_SSID="${WIFI_SSID}" \
	WIFI_PASSWORD="${WIFI_PASSWORD}" \
		pio run --environment ota_bootloader --target upload
