Feature: Updating existing "occurrences" sections

  Scenario: updating an existing occurrences section
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### what is it

      - foo
      - bar

      ### occurrences

      - [two](2.md)
      """
    And the workspace contains file "2.md" with content:
      """
      # Two

      ### what is it

      - an example of [one](1.md)
      """
    When fixing the TikiBase
    Then the workspace should contain the file "1.md" with content:
      """
      # One

      ### what is it

      - foo
      - bar

      ### occurrences

      - [Two (what is it)](2.md#what-is-it)
      """
    And file "2.md" is unchanged
