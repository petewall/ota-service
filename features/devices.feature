Feature: Devices
  Scenario: Listing with no devices
    Given the OTA service is running
    When I ask for the list of devices
    Then the request is successful
    And I receive an empty hash

  Scenario: Listing with one device
    Given the OTA service is running
    And an update request comes from DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0
    When I ask for the list of devices
    Then the request is successful
    And I receive a hash with 1 entry
    And it contains a device with mac DEVICE_MAC_ADDRESS running SAMPLE_FIRMWARE version 1.0.0

  Scenario: Listing with multiple devices
    Given the OTA service is running
    And an update request comes from DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    And an update request comes from DEVICE_2 running SAMPLE_FIRMWARE version 1.0.1
    When I ask for the list of devices
    Then the request is successful
    And I receive a hash with 2 entries
    And it contains a device with mac DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0
    And it contains a device with mac DEVICE_2 running SAMPLE_FIRMWARE version 1.0.1

  Scenario: Viewing devices on the registry page

  Scenario: A new discovered device updates the registry page

  Scenario: An updated device updates the registry page
