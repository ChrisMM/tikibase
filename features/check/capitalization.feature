Feature: Listing section types with different capitalization

  Scenario: the same section uses different capitalization
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### what is it

      ### how it works
      """
    And the workspace contains file "2.md" with content:
      """
      # Two

      ### What is it

      ### How it works
      """
    And the workspace contains file "3.md" with content:
      """
      # Three

      ### WHAT IS IT

      ### how it works
      """
    When checking the TikiBase
    Then it finds these sections with mixed capitalization:
      | how it works, How it works         |
      | what is it, What is it, WHAT IS IT |
