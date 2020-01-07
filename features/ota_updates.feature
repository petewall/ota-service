Feature: OTA updates
  Scenario: A device requests an update with no new version
    Given the OTA service is running
    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the service responds with no update

	Scenario: A device requests an update with a new version
    Given a firmware binary with type SAMPLE_FIRMWARE and version 1.0.1
    And the OTA service is running
    And there is a device DEVICE_MAC_ADDRESS with an assigned firmware type SAMPLE_FIRMWARE
    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    Then the service sends the firmware binary for SAMPLE_FIRMWARE with version 1.0.1

  Scenario: Updates show on the registry page
    Given a firmware binary with type SAMPLE_FIRMWARE and version 1.0.1
    And the OTA service is running
    And there is a device DEVICE_MAC_ADDRESS with an assigned firmware type SAMPLE_FIRMWARE
    When I view the registry page
    Then the registry page shows that the state of device DEVICE_MAC_ADDRESS is reassigned

    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    And I view the registry page
    Then the registry page shows that the state of device DEVICE_MAC_ADDRESS is updating

    When an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.1
    And I view the registry page
    Then the device list has a device with mac DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.1
    Then the registry page shows that the state of device DEVICE_MAC_ADDRESS is up to date
