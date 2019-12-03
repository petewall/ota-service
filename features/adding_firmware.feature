Feature: Adding firmware
  Scenario: Adding a firmware binary
    Given an empty firmware directory
    And the OTA service is running
    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive an empty list

    When a firmware binary for FEATURE_TEST_FIRMWARE with a version of 1.0.0 is added
    Then the service detects 1 binary

    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 1 entry
    And it contains a firmware for FEATURE_TEST_FIRMWARE with a version of 1.0.0
