Feature: Devices
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

  Scenario: Getting device properties
    Given the OTA service is running
    And an update request comes from DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    When I ask for the device properties for DEVICE_1
    Then the request is successful
    And I receive a hash
    And the result has a mac of DEVICE_1
    And the result has a firmwareType of SAMPLE_FIRMWARE
    And the result has a firmwareVersion of 1.0.0
    And the result has no assignedFirmware
    And the result has no deviceId

  Scenario: Setting device properties
    Given the OTA service is running
    And an update request comes from DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    When I assign a device id of Office to DEVICE_1
    Then the request is successful
    And the server updates DEVICE_1 with device id of Office

    When I ask for the device properties for DEVICE_1
    Then the request is successful
    And I receive a hash
    And the result has a mac of DEVICE_1
    And the result has an id of Office

  Scenario: Setting device id on the registry page
    Given the OTA service is running
    And an update request comes from DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0

    When I view the registry page
    And I enter a device id of Office to DEVICE_1 on the registry page
    Then the server updates DEVICE_1 with device id of Office

    When I view the registry page
    Then the device list has 1 entry
    And the device list has a device with mac DEVICE_1 with a device id of Office
