Feature: linkify the longest matching terms

    When linkifying a document
  I want that the longest matching term gets linkified
  So that the links are as precise as possible.

  Scenario: a document contains multiple matching terms
    Given the workspace contains file "1.md" with content:
      """
      # One

      Amazon makes Amazon Web Services
      """
    And the workspace contains file "amazon.md" with content:
      """
      # Amazon
      """
    And the workspace contains file "amazon-web-services.md" with content:
      """
      # Amazon Web Services
      """
    When linkifying
    Then the workspace should contain the file "1.md" with content:
      """
      # One

      [Amazon](amazon.md) makes [Amazon Web Services](amazon-web-services.md)
      """
