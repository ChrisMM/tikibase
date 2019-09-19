# Tikibase

[![Go Report Card](https://goreportcard.com/badge/github.com/kevgo/tikibase)](https://goreportcard.com/report/github.com/kevgo/tikibase)

_timeless wiki-like knowledge database_

Tikibase is an extremely robust knowledge base designed to be usable for
decades.

- semi-structured text storage to model knowledge
- links and backlinks
- storage format is human and machine readable and editable Markdown organized
  in files
- history is stored in Git
- offline availability via Git checkout
- simple CLI tools and bots to automate maintenance

### Why

Over the last couple of decades, I have used many types of software to store
things that I want to remember:

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

- it would be nice to reference other files via `#shortname` rather than
  creating a Markdown link to the file

### Development

- run a single unit test: `go test ./storage/foo_test.go`
