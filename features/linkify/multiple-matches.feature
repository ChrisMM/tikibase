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
      # Missing Links

      Amazon headquarter is not at the Amazon river.
      """
    When linkifying
    Then the workspace should contain the file "missing-links.md" with content:
      """
      # Missing Links

      [Amazon](amazon.md) headquarter is not at the Amazon river.
      """
