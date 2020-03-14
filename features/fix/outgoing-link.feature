Feature: Ignore outgoing links

  Scenario: a file contains an HTTP link
    Given the workspace contains file "1.md" with content:
      """
      # One
      [Google](http://google.com)

      <a href="http://google.com">Google again</a>
      """
    When fixing the TikiBase
    Then file "1.md" is unchanged
