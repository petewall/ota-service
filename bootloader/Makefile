.PHONY: build clean test

clean:
	rm -rf .pio

SRC_FILES := $(shell find . -path ./test -prune -false -o -name "*.cpp" -o -name "*.h")
TEST_FILES := $(shell find ./test -name "*.cpp" -o -name "*.h")

define invoke-pio
	FIRMWARE_VERSION="${FIRMWARE_VERSION}" \
	OTA_HOST="http://wallserver.local:8266" \
	WIFI_SSID="${WIFI_SSID}" \
	WIFI_PASSWORD="${WIFI_PASSWORD}" \
		pio $1 --environment ota_bootloader $2 $3
endef

test: build ${TEST_FILES}
	$(call invoke-pio,test,--verbose)

.pio/build/ota_bootloader/firmware.bin: platformio.ini ${SRC_FILES}
	$(call invoke-pio,test)

build: .pio/build/ota_bootloader/firmware.bin

check: ${SRC_FILES}
	$(call invoke-pio,check,--verbose,--skip-packages)

upload:
	$(call invoke-pio,run,--target,upload)
