Feature: don't linkify parts of URLs

  Scenario: match in URL
    Given the workspace contains file "byte.md" with content:
      """
      # Byte

      link: https://www.byte.com
      """
    When linkifying
    Then the workspace should contain the file "byte.md" with content:
      """
      # Byte

      link: https://www.byte.com
      """
