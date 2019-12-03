Feature: Listing Firmware
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
    Given a firmware binary with type FEATURE_TEST_FIRMWARE and version 1.0.0
    And the OTA service is running
    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 1 entry
    And it contains a firmware for FEATURE_TEST_FIRMWARE with a version of 1.0.0

  Scenario: Listing firmware on the registry page
    Given a firmware binary with type FIRMWARE_A and version 1.0.0
    Given a firmware binary with type FIRMWARE_C and version 1.0.0
    Given a firmware binary with type FIRMWARE_A and version 3.0.0
    Given a firmware binary with type FIRMWARE_B and version 1.0.0
    Given a firmware binary with type FIRMWARE_A and version 2.0.0
    And the OTA service is running
    When I view the registry page
    Then the firmware list has 5 entries

    And the firmware list contains a firmware for FIRMWARE_A with a version of 1.0.0
    And the firmware list contains a firmware for FIRMWARE_A with a version of 2.0.0
    And the firmware list contains a firmware for FIRMWARE_A with a version of 3.0.0
    And the firmware list contains a firmware for FIRMWARE_B with a version of 1.0.0
    And the firmware list contains a firmware for FIRMWARE_C with a version of 1.0.0
    And the firmware list is sorted by firmware then by version

  Scenario: Listing with multiple firmware types
    Given a firmware binary with type A_FIRMWARE and version 1.2.3
    And a firmware binary with type ANOTHER_FIRMWARE and version 4.5.6
    And the OTA service is running
    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 2 entries
    And it contains a firmware for A_FIRMWARE with a version of 1.2.3
    And it contains a firmware for ANOTHER_FIRMWARE with a version of 4.5.6

  Scenario: Listing with multiple firmware versions
    Given a firmware binary with type SAME_FIRMWARE and version 1.0.0
    And a firmware binary with type SAME_FIRMWARE and version 1.0.1
    And the OTA service is running
    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive a list with 2 entries
    And it contains a firmware for SAME_FIRMWARE with a version of 1.0.0
    And it contains a firmware for SAME_FIRMWARE with a version of 1.0.1
