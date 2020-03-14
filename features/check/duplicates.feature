Feature: Listing duplicates

  Scenario: duplicate sections
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### foo
      ### foo
      """
    And the workspace contains file "2.md" with content:
      """
      # Two

      ### bar
      ### bar
      """
    When checking the TikiBase
    Then it finds the duplicates:
      | 1.md#foo |
      | 2.md#bar |
