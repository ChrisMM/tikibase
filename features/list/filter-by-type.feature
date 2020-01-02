Feature: List items filtered by type

  Scenario: type in plain text
    Given the workspace contains file "markdown-it.md" with content:
      """
      # Markdown-It

      ### what is it
      - a Markdown parser
      - in JavaScript
      """
    When listing "parser"
    Then it finds:
      | Markdown-It |

  Scenario: type as link
    Given the workspace contains file "markdown-it.md" with content:
      """
      # Markdown-It

      ### what is it
      - a Markdown [parser](parser.md)
      - in JavaScript
      """
    When listing "parser"
    Then it finds:
      | Markdown-It |

  Scenario: different case
    Given the workspace contains file "markdown-it.md" with content:
      """
      # Markdown-It

      ### what is it
      - a Markdown parser
      - in JavaScript
      """
    When listing "markdown"
    Then it finds:
      | Markdown-It |
