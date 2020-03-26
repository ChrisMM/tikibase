Feature: Empty sections

  Scenario: empty document title
    Given the workspace contains file "1.md" with content:
      """
      #

      Hello
      """
    When checking the TikiBase
    Then it finds these documents with empty sections:
      | 1.md |

  Scenario: empty section
    Given the workspace contains file "1.md" with content:
      """
      # Foo

      ###

      Hello
      """
    When checking the TikiBase
    Then it finds these documents with empty sections:
      | 1.md |
