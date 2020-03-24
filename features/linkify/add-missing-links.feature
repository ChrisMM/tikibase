Feature: Add missing links

  Scenario: a file is missing a link
    Given the workspace contains file "amazon-web-services.md" with content:
      """
      # Amazon Web Services
      """
    And the workspace contains file "missing-links.md" with content:
      """
      # Missing Links

      Have you heard about Amazon Web Services?
      """
    When linkifying
    Then the workspace should contain the file "missing-links.md" with content:
      """
      # Missing Links

      Have you heard about [Amazon Web Services](amazon-web-services.md)?
      """

  Scenario: partial match
    Given the workspace contains file "byte.md" with content:
      """
      # Byte
      """
    And the workspace contains file "storage.md" with content:
      """
      # Storage

      several kilobytes of storage
      """
    When linkifying
    Then the workspace should contain the file "storage.md" with content:
      """
      # Storage

      several kilobytes of storage
      """
