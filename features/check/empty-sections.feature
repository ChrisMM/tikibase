Feature: Empty sections

  Scenario: empty document title
    Given the workspace contains file "1.md" with content:
      """
      #

      Hello
      """
    When checking the TikiBase
    Then it finds the empty document titles:
      | 1.md |
