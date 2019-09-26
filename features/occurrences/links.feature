Feature: Links

  Scenario: a link points to an anchor in a content section of another file
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
    When running Occurrences
    Then file "1.md" is unchanged
    And the workspace should contain the file "2.md" with content:
      """
      # Two
      ### benefits

      ### occurrences

      - [One (related)](1.md#related)
      """

  Scenario: a link points to an anchor in the title section of another file
    Given the workspace contains file "1.md" with content:
      """
      # One
      [benefits of two](2.md)
      """
    And the workspace contains file "2.md" with content:
      """
      # Two
      """
    When running Occurrences
    Then file "1.md" is unchanged
    And the workspace should contain the file "2.md" with content:
      """
      # Two

      ### occurrences

      - [One](1.md)
      """

  Scenario: two links to the same section
    Given the workspace contains file "1.md" with content:
      """
      # One
      ### related
      [benefits of two](2.md#benefits)
      [benefits of two](2.md#benefits)
      """
    And the workspace contains file "2.md" with content:
      """
      # Two
      ### benefits
      """
    When running Occurrences
    Then file "1.md" is unchanged
    And the workspace should contain the file "2.md" with content:
      """
      # Two
      ### benefits

      ### occurrences

      - [One (related)](1.md#related)
      """

  Scenario: a link points to an anchor within the same file
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### advantages
      See also the many [benefits](#benefits)
      """
    When running Occurrences
    Then file "1.md" is unchanged