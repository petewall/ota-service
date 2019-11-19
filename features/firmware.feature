Feature: Firmware
  Scenario: Listing with no firmware
    Given an empty firmware directory
    And the OTA service is running
    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive an empty list

  Scenario: Listing with no firmware on the registry page
    Given an empty firmware directory
    And the OTA service is running
    When I view the registry page
    Then the firmware list is empty

  Scenario: Listing with one firmware
    Given there is a firmware binary for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    And the OTA service is running
    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 1 entry
    And it contains a firmware for FEATURE_TEST_FIRMWARE with a version of 1.0.0

  Scenario: Listing with one firmware on the registry page
    Given there is a firmware binary for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    And the OTA service is running
    When I view the registry page
    Then the firmware list has 1 entry
    And the firmware list contains a firmware for FEATURE_TEST_FIRMWARE with a version of 1.0.0

  Scenario: Listing with multiple firmware types
    Given there is a firmware binary for A_FIRMWARE with a version of 1.2.3
    And there is a firmware binary for ANOTHER_FIRMWARE with a version of 4.5.6
    And the OTA service is running
    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 2 entries
    And it contains a firmware for A_FIRMWARE with a version of 1.2.3
    And it contains a firmware for ANOTHER_FIRMWARE with a version of 4.5.6

  Scenario: Listing with multiple firmware versions
    Given there is a firmware binary for SAME_FIRMWARE with a version of 1.0.0
    And there is a firmware binary for SAME_FIRMWARE with a version of 1.0.1
    And the OTA service is running
    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 2 entries
    And it contains a firmware for SAME_FIRMWARE with a version of 1.0.0
    And it contains a firmware for SAME_FIRMWARE with a version of 1.0.1

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
