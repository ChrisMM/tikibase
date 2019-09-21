Feature: ignore HTML links

  Scenario: a file contains an HTML link
    Given the workspace contains file "1.md" with content:
      """
      # One
      [Google](http://google.com)
      """
    When running Mentions
    Then file "1.md" is unchanged
