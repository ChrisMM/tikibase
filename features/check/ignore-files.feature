Feature: Ignoring files

  Scenario: Ignoring a resource
    Given the workspace contains a binary file "Makefile"
    And the workspace contains file "tikibase.yml" with content:
      """
      ignore:
        - Makefile
      """
    When checking the TikiBase
    Then it finds no errors
