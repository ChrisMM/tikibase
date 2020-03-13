<img src="tikibase.jpg" width="92" height="216" align="right">

# TikiBase Tools

[![CircleCI](https://circleci.com/gh/kevgo/tikibase.svg?style=shield)](https://circleci.com/gh/kevgo/tikibase)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevgo/tikibase)](https://goreportcard.com/report/github.com/kevgo/tikibase)

This repository provides automation to make working with a
[TikiBase](documentation/tikibase.md) more convenient.

### functionality

The `tikibase` tool in this repository provides these commands:

- **occurrences:** adds an `occurrences` section to each document that lists the
  documents that link to this document. This section only only contains
  documents that aren't otherwise mentioned in this document already.
- **list:** lists entries in this Tikibase
  - list all parsers: `tikibase list --is parser`
  - list all Markdown parsers: `tikibase list --is parser,markdown` or
    `tikibase list --is parser --is markdown`
- **check:** lists dead internal links in this Tikibase

You can run these commands using the CLI, on a CI server, or bundle them into a
bot to run them automatically on each change.

### development

- **make test:** run all tests
- **make help:** see all available tasks in the [Makefile](Makefile)

Install [scc](https://github.com/boyter/scc) to see code statistics.
