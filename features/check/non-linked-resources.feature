Feature: Finding non-linked resources

  Scenario: a resource without a link
    And the workspace contains a binary file "photo.jpg"
    And the workspace contains a binary file "contract.pdf"
    When checking the TikiBase
    Then it finds the non-linked resources:
      | contract.pdf |
      | photo.jpg    |
