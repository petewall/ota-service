#ifndef __WALLHOUSE_PERIODIC_ACTION_H__
#define __WALLHOUSE_PERIODIC_ACTION_H__

class PeriodicAction {
public:
  PeriodicAction(unsigned long interval);
  void check(unsigned long millis);
  virtual void run() = 0;

private:
  unsigned long interval;
  unsigned long next;
};

#endif // __WALLHOUSE_PERIODIC_ACTION_H__
