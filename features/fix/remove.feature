Feature: Removing obsolete "occurrences" sections

  Scenario: a file contains an obsolete "occurrences" section
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### what is it

      - foo
      - bar

      ### occurrences

      - [Two (what is it)](2.md#what-is-it)
      """
    And the workspace contains file "2.md" with content:
      """
      # Two

      ### what is it

      - an example of [one](1.md)

      ### occurrences

      Obsolete occurrences content.
      """
    When fixing the TikiBase
    Then file "1.md" is unchanged
    And the workspace should contain the file "2.md" with content:
      """
      # Two

      ### what is it

      - an example of [one](1.md)

      """
