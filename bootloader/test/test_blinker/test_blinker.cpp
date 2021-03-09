#include <Arduino.h>
#include <Blinker.h>
#include <unity.h>

const unsigned long TEST_BLINKER_INTERVAL = 100;

Blinker *blinker;

uint8_t i = 0;
const uint8_t max_blinks = 5;

void test_blinker_builtin_pin_number(void) {
  TEST_ASSERT_EQUAL(2, LED_BUILTIN);
}

void test_blinker_state_high(void) {
  TEST_ASSERT_EQUAL(HIGH, digitalRead(LED_BUILTIN));
  TEST_ASSERT_EQUAL(true, blinker->state);
}

void test_blinker_state_low(void) {
  TEST_ASSERT_EQUAL(LOW, digitalRead(LED_BUILTIN));
  TEST_ASSERT_EQUAL(false, blinker->state);
}

void setup() {
  delay(2000);
  UNITY_BEGIN();
  
  blinker = new Blinker(TEST_BLINKER_INTERVAL);
  TEST_MESSAGE("Blinker test suite initialized");
  RUN_TEST(test_blinker_builtin_pin_number);}


bool test_blinker_run() {
  if (i < max_blinks) {
    delay(TEST_BLINKER_INTERVAL / 2);

    blinker->check(millis());
    RUN_TEST(test_blinker_state_high);

    delay(TEST_BLINKER_INTERVAL);

    blinker->check(millis());
    RUN_TEST(test_blinker_state_low);

    delay(TEST_BLINKER_INTERVAL / 2);

    i++;
    return false;
  }

  TEST_MESSAGE("Blinker test suite done");
  return true;
}

void loop() {
  if (test_blinker_run()) {
    UNITY_END();
  }
}
