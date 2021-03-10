Feature: Assigning firmware types
    Scenario: Assigning a new firmware type to a device
    Given a firmware binary with type NEW_FIRMWARE and version 1.0.0
    And the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    When I assign a firmware type of NEW_FIRMWARE to DEVICE_MAC_ADDRESS
    Then the firmware NEW_FIRMWARE is assigned to DEVICE_MAC_ADDRESS

    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the request is successful
    And the service sends the firmware binary for NEW_FIRMWARE with version 1.0.0

    Scenario: Assigning a new firmware type to a device on the registry page
    Given a firmware binary with type NEW_FIRMWARE and version 1.0.0
    And the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    When I view the registry page
    And I select NEW_FIRMWARE from the dropdown of firmware for DEVICE_MAC_ADDRESS on the registry page
    Then the firmware NEW_FIRMWARE is assigned to DEVICE_MAC_ADDRESS

    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the request is successful
    And the service sends the firmware binary for NEW_FIRMWARE with version 1.0.0

    Scenario: Removing a firmware assignment
    Given a firmware binary with type NEW_FIRMWARE and version 1.0.0
    And the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 0.0.1
    When I assign a firmware type of NEW_FIRMWARE to DEVICE_MAC_ADDRESS
    Then the firmware NEW_FIRMWARE is assigned to DEVICE_MAC_ADDRESS

    When I view the registry page
    And I set the firmware type for DEVICE_MAC_ADDRESS to unassigned
    Then the device DEVICE_MAC_ADDRESS in unassigned

    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the service responds with no update
