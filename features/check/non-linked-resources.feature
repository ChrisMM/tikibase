Feature: Finding non-linked resources

  Scenario: a resource without a link
    Given the workspace contains file "1.md" with content:
      """
      # One

      [diagram](diagram.png)
      """
    And the workspace contains a binary file "photo.jpg"
    And the workspace contains a binary file "contract.pdf"
    And the workspace contains a binary file "diagram.png"
    When checking the TikiBase
    Then it finds the non-linked resources:
      | contract.pdf |
      | photo.jpg    |
