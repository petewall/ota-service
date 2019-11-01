Feature: Listing binaries
  Scenario: No binaries
    Given an empty binary directory
    And the OTA Service is running
    When I ask for the list of binaries
    Then the request is successful
    And I receive an empty hash

  Scenario: One binary
    Given a binary directory with one binary
    And the OTA Service is running
    When I ask for the list of binaries
    Then the request is successful
    And I receive a hash with a binary in a single device

  Scenario: Multiple binaries
    Given a binary directory with binaries for multiple devices
    And the OTA Service is running
    When I ask for the list of binaries
    Then the request is successful
    And I receive a hash with multiple devices
