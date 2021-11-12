#include <Arduino.h>
#include <Blinker.h>
#include <ESP8266httpUpdate.h>
#include <OTAClient.h>

Blinker* blinker;
OTAClient* ota;
ReliableNetwork *network;

#define HALF_SECOND 500
#define ONE_MINUTE 60 * 1000

// cppcheck-suppress unusedFunction
void setup() {
  Serial.begin(115200);
  blinker = new Blinker(HALF_SECOND);
  network = new ReliableNetwork(WIFI_SSID, WIFI_PASSWORD);
  network->connect();

  ota = new OTAClient(OTA_HOSTNAME, OTA_PORT, network, ONE_MINUTE);
}

// cppcheck-suppress unusedFunction
void loop() {
  unsigned long now = millis();
  blinker->check(now);
  network->check(now);
  ota->check(now);
}
