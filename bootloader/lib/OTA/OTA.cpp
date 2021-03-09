#include "OTA.h"
#include <Arduino.h>
#include <WiFiClient.h>
#include <ESP8266httpUpdate.h>

#define OTA_URL OTA_HOST "/api/update?firmware=" FIRMWARE_TYPE "&version=" FIRMWARE_VERSION

OTA::OTA(unsigned long interval)
: PeriodicAction(interval) {}

void OTA::run() {
  WiFiClient client;

  ESPhttpUpdate.setLedPin(LED_BUILTIN, LOW);
  Serial.println("[OTA] Sending request to " OTA_URL);
  t_httpUpdate_return result = ESPhttpUpdate.update(client, OTA_URL);
  switch(result) {
    case HTTP_UPDATE_FAILED:
      Serial.println("[OTA] Update failed");
      break;
    case HTTP_UPDATE_NO_UPDATES:
      Serial.println("[OTA] No updates");
      break;
    case HTTP_UPDATE_OK:
      Serial.println("[OTA] Updated!"); // may not be called since we reboot the ESP
      break;
  }
}