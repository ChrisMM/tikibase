Feature: don't linkify parts of existing links

  Scenario: existing link with match
    Given the workspace contains file "byte.md" with content:
      """
      # Byte
      """
    And the workspace contains file "storage.md" with content:
      """
      # Storage

      several [kilo-bytes](kilo-byte.md) of storage
      """
    When linkifying
    Then the workspace should contain the file "storage.md" with content:
      """
      # Storage

      several [kilo-bytes](kilo-byte.md) of storage
      """
