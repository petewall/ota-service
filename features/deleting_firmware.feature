Feature: Deleting firmware
  Scenario: Deleting a binary
    Given there is a firmware binary for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    And the OTA service is running
    When I send a delete request for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    Then the request is successful
    And the service detects 0 binaries

    When I ask for the list of firmware binaries
    Then the request is successful
    And I receive an empty list

  Scenario: Deleting a binary on the registry page
    Given there is a firmware binary for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    And the OTA service is running
    When I view the registry page
    And I click the delete button for FEATURE_TEST_FIRMWARE with a version of 1.0.0
    Then the service detects 0 binaries

    When I view the registry page
    Then the firmware list is empty
