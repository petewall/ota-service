#ifndef __OTA_SERVICE_BOOTLOADER_OTA_H__
#define __OTA_SERVICE_BOOTLOADER_OTA_H__

#include <PeriodicAction.h>

class OTA : public PeriodicAction {
public:
  explicit OTA(unsigned long interval);

protected:
  void run() override;
};

#endif // __OTA_SERVICE_BOOTLOADER_OTA_H__