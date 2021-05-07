.PHONY: build clean test

clean:
	rm -rf .pio

SRC_FILES := $(shell find . -path ./test -prune -false -o -name "*.cpp" -o -name "*.h")

~/.ssh/id_github:
	setup-github-ssh-key.sh

.pio/libdeps/depspulled: ~/.ssh/id_github
	pio lib install
	touch .pio/libdeps/depspulled

deps: .pio/libdeps/depspulled

define invoke-pio
	FIRMWARE_VERSION="${FIRMWARE_VERSION}" \
	OTA_HOSTNAME="http://wallserver.local" \
	OTA_PORT="8266" \
	WIFI_SSID="${WIFI_SSID}" \
	WIFI_PASSWORD="${WIFI_PASSWORD}" \
		pio $1 --environment ota-bootloader $2 $3
endef

.pio/build/ota-bootloader/firmware.bin: platformio.ini ${SRC_FILES}
	$(call invoke-pio,run)

build: .pio/libdeps/depspulled .pio/build/ota-bootloader/firmware.bin

check: ${SRC_FILES} .pio/libdeps/depspulled
	$(call invoke-pio,check,--verbose,--skip-packages)

upload:
	$(call invoke-pio,run,--target,upload)
