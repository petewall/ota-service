Feature: Devices
  Scenario: Listing with no devices
    Given the OTA service is running
    When I ask for the list of devices
    Then the request is successful
    And I receive an empty hash

  Scenario: Listing with no devices on the registry page
    Given the OTA service is running
    When I view the registry page
    Then the device list is empty

  Scenario: Listing with one device
    Given the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    When I ask for the list of devices
    Then the request is successful
    And I receive a hash with 1 entry
    And it contains a device with mac DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0

  Scenario: Listing with one device on the registry page
    Given the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    When I view the registry page
    Then the device list has 1 entry
    And the device list has a device with mac DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0

  Scenario: Listing with multiple devices
    Given the OTA service is running
    And an update request comes from DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    And an update request comes from DEVICE_2 running SAMPLE_FIRMWARE version 1.0.1
    When I ask for the list of devices
    Then the request is successful
    And I receive a hash with 2 entries
    And it contains a device with mac DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    And it contains a device with mac DEVICE_2 running SAMPLE_FIRMWARE version 1.0.1

  Scenario: Listing with multiple devices on the registry page
    Given the OTA service is running
    And an update request comes from DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    And an update request comes from DEVICE_2 running SAMPLE_FIRMWARE version 1.0.1
    When I view the registry page
    Then the device list has 2 entries
    And the device list has a device with mac DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    And the device list has a device with mac DEVICE_2 running SAMPLE_FIRMWARE version 1.0.1
