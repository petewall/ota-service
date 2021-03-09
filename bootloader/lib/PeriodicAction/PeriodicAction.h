#ifndef __OTA_SERVICE_BOOTLOADER_PERIODIC_ACTION_H__
#define __OTA_SERVICE_BOOTLOADER_PERIODIC_ACTION_H__

class PeriodicAction {
public:
  explicit PeriodicAction(unsigned long interval);
  void check(unsigned long millis);
  virtual void run() = 0;

private:
  unsigned long interval;
  unsigned long next;
};

#endif // __OTA_SERVICE_BOOTLOADER_PERIODIC_ACTION_H__
