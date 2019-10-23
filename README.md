<img src="tikibase.jpg" width="92" height="216" align="right">

# TikiBase

[![CircleCI](https://circleci.com/gh/kevgo/tikibase.svg?style=shield)](https://circleci.com/gh/kevgo/tikibase)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevgo/tikibase)](https://goreportcard.com/report/github.com/kevgo/tikibase)

TikiBase is a timeless, low-tech wiki-like knowledge base designed to last for
decades. It makes semi-structured rich-text information available on all current
and future compute platforms.

### how it works

All forms of compute are platform-specific. Platforms change every decade. To
remain usable over decades, TikiBase is pure human-accessible data: Markdown
documents stored in a folder. All you need to use TikiBase is a Markdown or text
editor. TikiBase provides lightweight tooling (this codebase) for automation of
peripheral automation like managing hyperlinks.

### why

Over the last couple of decades, I have used a multitude of software products
for Windows, Linux, MacOS, or the web to store information that I want to
remember:

- **outliners:** OmniOutliner, Workflowy, Dynalist, ThinkLinkr, etc
- **note-taking applications:** Evernote, SimpleNote, Apple Notes, Notational
  Velocity
- **todo-lists:** Todoist, TickTick, Clear, etc
- multiple **custom-built solutions** that I didn't have time to develop into a
  full product

All these solutions have been lacking functionality or became unsupported at
some point. Or I changed platforms and couldn't use the applications anymore.
Whatever fancy new web/mobile/wearable tool shows up next, I can already feel
how it also will look and feel outdated in no more than 10 years and will be
almost unusable in 15, on at least some platforms that I'll use at that time. As
an example, try using OrgMode on mobile!

### functionality

You use (create, edit, search) a TikiBase completely manually, for example using
the GitHub or GitLab UI, or your favorite MarkDown editor. On top of that, the
TikiBase tool in this repository provides these commands:

- **occurrences:** adds an `occurrences` section to each document that lists the
  documents that link to this document. This section only only contains
  documents that aren't otherwise mentioned in this document already.

You can run these commands using the CLI, on a CI server, or bundle them into a
bot to run them automatically on each change.

### best practices

- **avoid deep hierarchies:** Taxonomies are highly debatable, even between your
  current and future self. Keep information organized in a relatively flat
  namespace and retrieve information alphabetically or via fulltext search.

### limitations

- it would be nice to reference other files via **#shortname** rather than
  having to create a Markdown link to the file

### development

- **make test:** run all tests
- see the [Makefile](Makefile) for available commands
