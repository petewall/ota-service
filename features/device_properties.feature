Feature: Device properties
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

  Scenario: Getting a single property
    Given the OTA service is running
    And an update request comes from DEVICE_1 running SAMPLE_FIRMWARE version 1.0.0

    When I ask for the mac property for DEVICE_1
    Then the request is successful
    And I receive the value DEVICE_1

    When I ask for the firmwareType property for DEVICE_1
    Then the request is successful
    And I receive the value SAMPLE_FIRMWARE

    When I ask for the firmwareVersion property for DEVICE_1
    Then the request is successful
    And I receive the value 1.0.0

    When I ask for the assignedFirmware property for DEVICE_1
    Then the request returns no content

    When I ask for the id property for DEVICE_1
    Then the request returns no content

    When I ask for the foo property for DEVICE_1
    Then the request returns not found
