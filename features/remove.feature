Feature: Removing obsolete "mentions" sections

  Scenario: a file contains an obsolete "mentions" section
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### what is it

      - foo
      - bar

      ### mentions

      - [Two](2.md#what-is-it)
      """
    And the workspace contains file "2.md" with content:
      """
      # Two

      ### what is it

      - an example of [one](1.md)

      ### mentions

      Obsolete mentions content.
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
    And the workspace should contain the file "2.md" with content:
      """
      # Two

      ### what is it

      - an example of [one](1.md)

      """
