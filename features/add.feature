Feature: Adding mentions

  Scenario: a file contains no "mentions" section
    Given the workspace contains file "1.md" with content:
      """
      # One

      ## what is it

      - foo
      - bar

      ## what does it

      - it foos
      - it bars
      """
    And the workspace contains file "2.md" with content:
      """
      # Two

      ## what is it

      - an example of [1.md](one)
      """
    When running Mentions
    Then the workspace should contain the file "1.md" with content:
      """
      # One

      ## what is it

      - foo
      - bar

      ## what does it

      - it foos
      - it bars

      ## mentions

      - [2.md](Two)
      """
