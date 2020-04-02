Feature: ignore blockquotes

  When adding missing hyperlinks
  I want that text inside blockquotes is ignored
  So what hyperlinks don't mess up source code.

  Scenario: block quotes contain link targets
    Given the workspace contains file "authorization.md" with content:
      """
      # authorization
      """
    And the workspace contains file "basic-authentication.md" with content:
      """
      # basic authentication

      ```
      Authorization: Basic ZGVtbzpwQDU1dzByZA==
      ```
      """
    When linkifying
    Then the workspace should contain the file "basic-authentication.md" with content:
      """
      # basic authentication

      ```
      Authorization: Basic ZGVtbzpwQDU1dzByZA==
      ```
      """
