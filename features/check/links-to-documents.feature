Feature: Checking internal links to documents

  Scenario Outline: working links
    Given the workspace contains file "1.md" with content:
      """
      # One
      <DESC>: <LINK>
      ### Beta
      """
    And the workspace contains file "2.md" with content:
      """
      # Two
      ### Hello
      """
    When checking the TikiBase
    Then it finds no errors

    Examples:
      | DESC                                    | LINK                         |
      | Markdown link to file                   | [Two](2.md)                  |
      | HTML link to file                       | <a href="2.md">Two</a>       |
      | Markdown link to anchor in another file | [Two](2.md#hello)            |
      | HTML link to anchor in another file     | <a href="2.md#hello">Two</a> |
      | Markdown link to anchor in same file    | [Beta](#beta)                |
      | HTML link to anchor in same file        | <a href="#beta">Beta</a>     |

  Scenario Outline: broken links
    Given the workspace contains file "1.md" with content:
      """
      # One
      <DESC>: <LINK>
      """
    And the workspace contains file "2.md" with content:
      """
      # Two
      """
    When checking the TikiBase
    Then it finds the broken links:
      | FILE | LINK            |
      | 1.md | <REPORTED HREF> |

    Examples:
      | DESC                                         | LINK                        | REPORTED HREF |
      | MD link to missing file                      | [Three](3.md)               | 3.md          |
      | HTML link to missing file                    | <a href="3.md">Three</a>    | 3.md          |
      | MD link to missing anchor in existing file   | [Two](2.md#zonk)            | 2.md#zonk     |
      | HTML link to missing anchor in existing file | <a href="2.md#zonk">Two</a> | 2.md#zonk     |
      | MD link to missing anchor in same file       | [One](1.md#zonk)            | 1.md#zonk     |
      | HTML link to missing anchor in same file     | <a href="1.md#zonk">One</a> | 1.md#zonk     |
