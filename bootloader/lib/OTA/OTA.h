#ifndef __WALLHOUSE_OTA_H__
#define __WALLHOUSE_OTA_H__

#include <PeriodicAction.h>

class OTA : public PeriodicAction {
public:
  OTA(unsigned long interval);
  void run();
};

#endif // __WALLHOUSE_OTA_H__