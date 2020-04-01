Feature: multiple matches

  When managing a TikiBase
  I want that only the first match is linkified
  So that my document isn't plastered with too many annoying links.

  Scenario: multiple matches
    Given the workspace contains file "amazon.md" with content:
      """
      # Amazon
      """
    And the workspace contains file "missing-links.md" with content:
      """
      # missing links

      Amazon headquarter is not at the Amazon river.
      """
    When linkifying
    Then the workspace should contain the file "missing-links.md" with content:
      """
      # missing links

      [Amazon](amazon.md) headquarter is not at the Amazon river.
      """

  Scenario: multiple matches and partial match
    Given the workspace contains file "user-stories.md" with content:
      """
      # user stories
      """
    And the workspace contains file "security-user-stories.md" with content:
      """
      # security user stories
      """
    And the workspace contains file "missing-links.md" with content:
      """
      # missing links

      I already have a link here: [security user stories](security-user-stories.md)
      and don't want another one here: security user stories
      """
    When linkifying
    Then the workspace should contain the file "missing-links.md" with content:
      """
      # missing links

      I already have a link here: [security user stories](security-user-stories.md)
      and don't want another one here: security user stories
      """
