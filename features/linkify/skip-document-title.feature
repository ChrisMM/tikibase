Feature: don't linkify document title

  When linkifying
  I don't want that all links occur only in the document text
  So that titles remain simple non-rich text

  Scenario: match in document title
    Given the workspace contains file "byte.md" with content:
      """
      # byte
      """
    And the workspace contains file "kilo-byte.md" with content:
      """
      # Kilo Byte

      roughly a thousand times more
      """
    When linkifying
    Then the workspace should contain the file "kilo-byte.md" with content:
      """
      # Kilo Byte

      roughly a thousand times more
      """
