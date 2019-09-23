# TikiBase

[![CircleCI](https://circleci.com/gh/kevgo/tikibase.svg?style=shield)](https://circleci.com/gh/kevgo/tikibase)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevgo/tikibase)](https://goreportcard.com/report/github.com/kevgo/tikibase)

### what is it

- a timeless wiki-like knowledge database
- a robust knowledge base designed

### what does it

- makes various forms of information usable for decades, on all current and
  future compute platforms

### how it works

- since all compute is specific to platforms or at least computing paradigms,
  TikiBase is pure data
- TikiBase is completely manually usable, with some optional tooling for
  lightweight automation

data store:

- human and machine-editable information store: a folder containing Markdown
  files
- semi-structured document format: documents contain semantically meaningful
  sections
- logical connection between documents via links and backlinks

### Why

Over the last couple of decades, I have used many types of software products to
store information that I want to remember:

- various outliners on Windows, MacOS, and the web
- Evernote
- Simplenote
- Apple Notes
- Dynalist
- Notational Velocity
- various custom-built web and desktop applications

All of them have been limited and/or became unsupported at some point. Or I
changed platforms and couldn't use them anymore.

- off-the-shelf tools lack structure for knowledge management and backlinks.

### Limitations

- it would be nice to reference other files via **#shortname** rather than
  having to create a Markdown link to the file

### Development

- run a single unit test: `go test ./storage/foo_test.go`
