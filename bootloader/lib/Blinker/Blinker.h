#pragma once

#include <PeriodicAction.h>

class Blinker : public PeriodicAction {
public:
  explicit Blinker(unsigned long interval);

protected:
  bool run() override;

private:
  bool state;
  friend void test_blinker_state_high();
  friend void test_blinker_state_low();
};
