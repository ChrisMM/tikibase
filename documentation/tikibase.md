# TikiBase

A successful knowledge management solution must live for many decades, have the
ability to rejuvenate itself (modernize outdated content, improve the
organization of things), connect existing ideas in new ways, and capture ideas
that don’t fit into or contradict the existing system.

A TikiBase is such a knowledge base. It is a timeless, low-tech, wiki-like
knowledge base designed to last for decades. It makes semi-structured rich-text
information embedded with
[RDF](https://en.wikipedia.org/wiki/Resource_Description_Framework)- like
associations available on all current and future compute platforms. It can be
used completely manually.

- timeless knowlegde base
- extremely robust knowledge base
- low-tech knowledge base

### what does it

- stores everything I know
- lasts for decades
- allows gradual refinement of ideas by adding details while maintaining a
  high-level overview
- makes semi-structured rich-text information available in a nice format on all
  current and future compute platforms

### how it works

- the database consists of Markdown documents stored in a folder
  - Markdown is long-lived, relatively rich, structured, and machine +
    human-readable
  - the filename is the unique ID in simple textual form (all lowercase, no
    spaces and special characters)
  - the note title is the human-readable form (with proper capitalization,
    spaces, and special characters)
  - sections are document parts that have a header
  - files don't need to have a strict order of H1, H2, H3, etc. One can use only
    H2 and H4 and it would still recognize them as level 1 and 2 of the
    document.
- you edit the database content directly using a text or Markdown editor
- lineage and versioning via Git
- lightweight, portable tooling (this codebase) in CLI or Bot form automates
  peripheral management tasks to keep the data store in a good shape
  - features:
    - formatting
    - checking hyperlinks and back-references: find dead or missing links
    - populating lists of links to other notes: "mentions"
    - find old/outdated content
  - tooling is written in Go because of timelessness (Go is simple to use, its
    binaries don't need anything to run)

### why

- durability
  - any form of sophisticated database or binary format will become obsolete,
    even open-source
- platform independence
  - all forms of compute are platform-specific
  - platforms change every decade
  - I change platforms even faster: Windows, Mac, Linux, ChromeOS
  - today I use multiple platforms: phone, PC
- specific functionality
  - off-the-shelf solutions don't have all the features I need
  - plain text is too simple and not machine-parsable
- simplicity
  - any real product, even a web-based one, requires too much development work
    and polish to be useful.
  - I need to work directly with the raw database.
  - it needs to be simple enough that good working solutions can be prototyped.

### competition

- **outliners:** OmniOutliner, Workflowy, Dynalist, ThinkLinkr, etc
- **note-taking applications:** Evernote, SimpleNote, Apple Notes, Notational
  Velocity
- **todo-lists:** Todoist, TickTick, Clear, etc
- multiple **custom-built solutions** that I didn't have time to develop into a
  full product

### best practices

- **avoid deep hierarchies:** Taxonomies are highly debatable, even between your
  current and future self. Keep information organized in a relatively flat
  namespace and retrieve information alphabetically or via fulltext search.
- metadata should occur naturally in the document content
  - status (idea, draft, verified etc) is in a "status" section
  - keywords (programming, DevOps) should be naturally in the content,
    especially the "what is it" and "what does it" sections
- there are folders for the different types of information (contact, concept,
  how-to)
- information organization: there should be only one way to organize information
  - **file:** a top-level concept that can stand on its own, each file contains
    all the information for its concept
  - **section:** separate dimension of looking at something ("what is it", "what
    does it", "why we need it", "how it works", "status", "related",
    "competition", "vendors", "advantages", "challenges", "Q&A"
  - **bullet point:** data within a dimension

### limitations

- it would be nice to reference other files via **#shortname** rather than
  having to create a Markdown link to the file. But that would reduce
  readability.

### Q&A

- **How is this different from a wiki?** The information is in Wiki format. But
  this doesn’t depend on a particular Wiki product that will be obsolete in a
  short while
