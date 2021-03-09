#ifndef __OTA_SERVICE_BOOTLOADER_BLINKER_H__
#define __OTA_SERVICE_BOOTLOADER_BLINKER_H__

#include <PeriodicAction.h>

class Blinker : public PeriodicAction {
public:
  Blinker(unsigned long interval);
  void run();

private:
  bool state;
  friend void test_blinker_state_high();
  friend void test_blinker_state_low();
};

#endif // __OTA_SERVICE_BOOTLOADER_BLINKER_H__