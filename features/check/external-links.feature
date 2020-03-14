Feature: Ignore external links

  Scenario Outline: broken external link
    Given the workspace contains file "1.md" with content:
      """
      # One

      <LINK>
      """
    When checking the TikiBase
    Then it finds no errors

    Examples:
      | LINK                                                           |
      | [markdown HTTP](https://zonkountaunthaeuntoheunth.com)         |
      | [markdown HTTPS](https://zonkountaunthaeuntoheunth.com)        |
      | <a href="http://zonkountaunthaeuntoheunth.com">HTML HTTP</a>   |
      | <a href="https://zonkountaunthaeuntoheunth.com">HTML HTTPS</a> |
