#include "PeriodicAction.h"

PeriodicAction::PeriodicAction(unsigned long interval)
: interval(interval), next(0)
{}

void PeriodicAction::check(unsigned long millis) {
  if (millis > this->next) {
    this->next = millis + this->interval;
    this->run();
  }
}
