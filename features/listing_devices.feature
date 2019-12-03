Feature: Listing devices
  Scenario: Listing with no devices
    Given the OTA service is running
    When I ask for the list of devices
    Then the request is successful
    And I receive an empty list

  Scenario: Listing with no devices on the registry page
    Given the OTA service is running
    When I view the registry page
    Then the device list is empty

  Scenario: Listing devices on the registry page
    Given the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS_2 running SAMPLE_FIRMWARE version 1.0.0
    And an update request comes from DEVICE_MAC_ADDRESS_1 running SAMPLE_FIRMWARE version 1.0.0
    When I view the registry page
    Then the device list has 2 entries
    And the device list has a device with mac DEVICE_MAC_ADDRESS_1 running SAMPLE_FIRMWARE version 1.0.0
    And the device list has a device with mac DEVICE_MAC_ADDRESS_2 running SAMPLE_FIRMWARE version 1.0.0
    And the device list is sorted by mac

  Scenario: Listing devices
    Given the OTA service is running
    And an update request comes from DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    And an update request comes from DEVICE_2 running SAMPLE_FIRMWARE version 1.0.1
    When I ask for the list of devices
    Then the request is successful
    And I receive a list with 2 entries
    And it contains a device with mac DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    And it contains a device with mac DEVICE_2 running SAMPLE_FIRMWARE version 1.0.1
