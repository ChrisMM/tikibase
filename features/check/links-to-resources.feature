Feature: Checking links to internal resources

  Scenario Outline: link to existing resource
    Given the workspace contains file "1.md" with content:
      """
      # One
      <DESC>: <LINK>
      """
    And the workspace contains a binary file "photo.jpg"
    When checking the TikiBase
    Then it finds no errors

    Examples:
      | DESC      | LINK                          |
      | MD link   | [photo](photo.jpg)            |
      | HTML link | <a href="photo.jpg">photo</a> |

  Scenario Outline: link to non-existing resource
    Given the workspace contains file "1.md" with content:
      """
      # One
      <DESC>: <LINK>
      """
    When checking the TikiBase
    Then it finds the broken links:
      | FILE | LINK      |
      | 1.md | photo.jpg |

    Examples:
      | DESC      | LINK                          |
      | MD link   | [photo](photo.jpg)            |
      | HTML link | <a href="photo.jpg">photo</a> |
