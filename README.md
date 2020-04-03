<img src="tikibase.jpg" width="92" height="216" align="right">

# Tikibase Tools

[![CircleCI](https://circleci.com/gh/kevgo/tikibase.svg?style=shield)](https://circleci.com/gh/kevgo/tikibase)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevgo/tikibase)](https://goreportcard.com/report/github.com/kevgo/tikibase)

The un-database.

The `tikibase` tool provides these commands:

- **check:** verifies the consistency of this Tikibase:
  - broken internal links
  - resources (non-markdown files) that aren't linked to from a markdown file
  - multiple occurrences of the same section in a file
  - inconsistent capitalization of sections
  - documents with empty titles
- **find:** semantic search
  - find all parsers: `tikibase find --is parser`
  - find all Markdown parsers: `tikibase find --is parser,markdown` or
    `tikibase find --is parser --is markdown`
- **fix:** fixes all auto-fixable issues
  - adds `occurrences` sections containing unmentioned backlinks
- **linkify:** adds missing links to this Tikibase.
- **pitstop:** runs all fixes and checks. Run this regularly while making
  changes to a Tikibase. Use the alias `ps` to type less.
- **stats:** shows statistics about this Tikibase
- **version:** shows the version of the installed tool

### configuration

Update the file **tikibase.yml** in your Tikibase directory to configure this
tool. Here is the default configuration:

```yml
ignore:
  - node_modules
  - vendor
```

This supports simple path matching for now, please
[reach out](https://github.com/kevgo/tikibase/issues/new) if you need more
sophisticated globbing.

### development

- **make ps:** run after making changes to the code base
- **make help:** see all available tasks in the [Makefile](Makefile)

Install [scc](https://github.com/boyter/scc) to see code statistics via
`make stats`.
