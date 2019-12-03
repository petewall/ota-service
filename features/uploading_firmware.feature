Feature: Uploading firmware
  Scenario: Uploading a brand new firmware binary
    Given an empty firmware directory
    And the OTA service is running
    When I send a binary file for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    Then the request is successful
    And the a binary file for FEATURE_TEST_FIRMWARE with a version of 1.0.0 exists in the firmware directory
    And the service detects 1 binary

    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 1 entry
    And it contains a firmware for FEATURE_TEST_FIRMWARE with a version of 1.0.0

  Scenario: Uploading a new version of an exisitng firmware binary
    Given there is a firmware binary for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    And the OTA service is running
    When I send a binary file for FEATURE_TEST_FIRMWARE with a version of 2.0.0
    Then the request is successful
    And the a binary file for FEATURE_TEST_FIRMWARE with a version of 2.0.0 exists in the firmware directory
    And the service detects 2 binaries

    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 2 entries
    And it contains a firmware for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    And it contains a firmware for FEATURE_TEST_FIRMWARE with a version of 2.0.0

  Scenario: Overwriting an existing binary
    Given there is a firmware binary for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    And the OTA service is running
    When I send a binary file for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    Then the request is successful
    And the a binary file for FEATURE_TEST_FIRMWARE with a version of 1.0.0 exists in the firmware directory
    And the service detects 1 binary

    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 1 entry
    And it contains a firmware for FEATURE_TEST_FIRMWARE with a version of 1.0.0
