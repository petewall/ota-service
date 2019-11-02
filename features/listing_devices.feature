Feature: Listing devices
  Scenario: No devices
    Given the OTA Service is running
    When I ask for the list of devices
    Then the request returns a 200
    And I receive an empty hash

  Scenario: New device
    Given the OTA Service is running
    When I ask for a binary with a current version of 1.0.0
    And I ask for the list of devices
    Then the request returns a 200
    And I receive a hash with the device with a version of 1.0.0

  Scenario: Updated device
    Given there is a binary with a version of 1.0.1
    And the OTA Service is running
    When I ask for a binary with a current version of 1.0.0
    And I ask for the list of devices
    Then the request returns a 200
    And I receive a hash with the device with a version of 1.0.1
