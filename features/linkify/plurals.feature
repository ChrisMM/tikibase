Feature: consider pluralized forms

  Scenario: a documents contains the pluralized title of another document
    Given the workspace contains file "byte.md" with content:
      """
      # byte
      """
    And the workspace contains file "kilobyte.md" with content:
      """
      # kilobyte

      A thousand bytes.
      """
    When linkifying
    Then the workspace should contain the file "kilobyte.md" with content:
      """
      # kilobyte

      A thousand [bytes](byte.md).
      """
