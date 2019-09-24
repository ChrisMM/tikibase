Feature: Link to anchor

  Scenario: a link points to an anchor within another file
    Given the workspace contains file "1.md" with content:
      """
      # One
      ### related
      [benefits of two](2.md#benefits)
      """
    And the workspace contains file "2.md" with content:
      """
      # Two
      ### benefits
      """
    When running Mentions
    Then file "1.md" is unchanged
    And the workspace should contain the file "2.md" with content:
      """
      # Two
      ### benefits

      ### mentions

      - [One (related)](1.md#related)
      """
