#ifndef __OTA_SERVICE_BOOTLOADER_OTA_H__
#define __OTA_SERVICE_BOOTLOADER_OTA_H__

#include <PeriodicAction.h>

class OTA : public PeriodicAction {
public:
  OTA(unsigned long interval);
  void run();
};

#endif // __OTA_SERVICE_BOOTLOADER_OTA_H__