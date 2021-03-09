#include "Blinker.h"
#include <Arduino.h>

Blinker::Blinker(unsigned long interval)
: PeriodicAction(interval) {
  state = false;
  pinMode(LED_BUILTIN, OUTPUT);
}

void Blinker::run() {
  this->state = !this->state;
  digitalWrite(LED_BUILTIN, this->state ? HIGH : LOW);
  Serial.print("[Blinker] ");
  Serial.println(this->state ? "on" : "off");
}