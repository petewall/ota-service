Feature: Updates
  Scenario: A device requests an update with no new version
    Given the OTA service is running
    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the service responds with no update

	Scenario: A device requests an update with a new version
    Given there is a firmware binary for SAMPLE_FIRMWARE with a version of 1.0.1
    And the OTA service is running
    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the service sends the firmware binary for SAMPLE_FIRMWARE with version 1.0.1

	Scenario: Assigning a new firmware type to a device
    Given there is a firmware binary for NEW_FIRMWARE with a version of 1.0.0
    And the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    When I assign a firmware type of NEW_FIRMWARE to DEVICE_MAC_ADDRESS
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the request is successful
    Then the service sends the firmware binary for NEW_FIRMWARE with version 1.0.0

	Scenario: Assigning a new firmware type to a device on the registry page
    Given there is a firmware binary for NEW_FIRMWARE with a version of 1.0.0
    And the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    When I view the registry page
    And I select NEW_FIRMWARE from the dropdown of firmware for DEVICE_MAC_ADDRESS on the registry page
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the request is successful
    And the service sends the firmware binary for NEW_FIRMWARE with version 1.0.0

  Scenario: Updates show on the registry page
    Given there is a firmware binary for SAMPLE_FIRMWARE with a version of 1.0.1
    And the OTA service is running

    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    And I view the registry page
    Then the registry page shows that the state of device DEVICE_MAC_ADDRESS is updating

    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.1
    And I view the registry page
    Then the device list has a device with mac DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.1
