Feature: Ignoring hidden files

  Scenario: hidden resource
    Given the workspace contains a binary file ".prettierrc"
    When checking the TikiBase
    Then it finds no errors
