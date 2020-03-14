Feature: Ignore non-Markdown files

  Scenario: a TikiBase contains images
    Given the workspace contains file "1.md" with content:
      """
      # One
      """
    And the workspace contains file "2.md" with content:
      """
      # Two
      ### what is it
      - an example of [one](1.md)
      """
    And the workspace contains a binary file "1.png"
    When fixing the TikiBase
    Then the workspace should contain the file "1.md" with content:
      """
      # One

      ### occurrences

      - [Two (what is it)](2.md#what-is-it)
      """
    And file "2.md" is unchanged

  Scenario: a link points to an image file
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### advantages
      ![benefits](benefits.jpg)
      """
    When fixing the TikiBase
    Then file "1.md" is unchanged
