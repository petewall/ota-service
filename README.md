# OTA Service

## API

* `GET /update?firmware=<firmwareType>&version=<firmwareVersion>` - Process firmware update request
  * Parameters:
    * `firmware` - The firmware type of the requesting device.
    * `version` - The firmware version of the requesting device.
  * Required headers:
    * `x-esp8266-sta-mac` - The MAC address of the requesting device.
  * Return options:
    * `200 OK` - Content will also contain the binary of the updated firmware.
    * `304 Not Modified` - Indicates that there are no new firmware for this device.
    * `400 Bad Request` - Indicates missing MAC or current firmware.
    * `500 Internal Server Error` - Indicates an error processing the upgrade. See the message for more information. 

## Upgrade algorithm

* Get the device record from the [device service](https://github.com/petewall/device-service)
  * If the device is yet unknown, record it and return 304
* If a specific firmware type and version are assigned:
  * If the device is currently using that version, return 304
  * If the device is using some other firmware, return the assigned firmware
* If a firmware type is assigned, but no version:
  * Get the latest firmware of that type from the [firmware service](https://github.com/petewall/firmware-service)
  * If the latest is newer than what's currently on the device, return that firmware
  * If the latest is the same or older than what's currently on the device, return 304 
* If the device has no assigned firmware, return 304
