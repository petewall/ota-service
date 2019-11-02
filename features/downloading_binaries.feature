Feature: Downloading binaries
  Scenario: Unknown device
    Given the OTA Service is running
    When I ask for a binary for an unknown device
    Then the request returns a 304

  Scenario: No newer binary
    Given there is a binary with a version of 1.0.0
    And the OTA Service is running
    When I ask for a binary with a current version of 1.0.0
    Then the request returns a 304

  Scenario: Newer binary
    Given there is a binary with a version of 1.0.1
    And the OTA Service is running
    When I ask for a binary with a current version of 1.0.0
    Then the request returns a 200
    And the binary file is sent
