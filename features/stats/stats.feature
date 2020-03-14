Feature: Statistics

  Scenario: normal codebase
    Given the workspace contains file "1.md" with content:
      """
      # One

      ### section 1

      - [link 1](1.md)
      - <a href="1.md">link 2</a>

      ### section 2
      """
    And the workspace contains file "2.md" with content:
      """
      # Two

      ### section 3

      - [link 3](2.md)
      - <a href="2.md">link 4</a>

      ### section 4
      """
    And the workspace contains a binary file "resource1.jpg"
    And the workspace contains a binary file "resource2.pdf"
    When running Statistics
    Then it provides the statistics:
      | DOCUMENTS | 2 |
      | SECTIONS  | 4 |
      | LINKS     | 4 |
      | RESOURCES | 2 |
    And it finds the section types:
      | section 1 |
      | section 2 |
      | section 3 |
      | section 4 |
