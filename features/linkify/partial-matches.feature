Feature: don't linkify partial matches

  When linkifying
  I don't want that partial word matches get linkified
  So that the links are semantically correct


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
