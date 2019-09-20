Feature: Updating existing "mentions" sections

  Scenario: updating an existing mentions section
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### what is it

      - foo
      - bar

      ### mentions

      Old content.
      """
    And the workspace contains file "2.md" with content:
      """
      # Two

      ### what is it

      - an example of [one](1.md)
      """
    When running Mentions
    Then the workspace should contain the file "1.md" with content:
      """
      # One

      ### what is it

      - foo
      - bar

      ### mentions

      - [Two](2.md#what-is-it)
      """
    And file "2.md" is unchanged
