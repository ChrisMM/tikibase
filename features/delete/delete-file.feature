Feature: deleting a file

  Scenario: deleting a file with links
    Given the workspace contains file "old.md" with content:
      """
      # Old
      """
    And the workspace contains file "1.md" with content:
      """
      # One

      This feeld [old](old.md).
      """
    When deleting file "old.md"
    Then the workspace should no longer contain file "old.md"
    And the workspace should contain the file "1.md" with content:
      """
      # One

      This feeld old.
      """
