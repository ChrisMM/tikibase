# TikiBase

A TikiBase is a timeless, low-tech, wiki-like knowledge base designed to last
for decades. It makes semi-structured rich-text information available on all
current and future compute platforms. It can be used completely manually.

### how it works

You use (create, edit, search) a TikiBase completely manually, for example using
the GitHub or GitLab UI, or your favorite MarkDown editor.

All forms of compute are platform-specific. Platforms change every decade.

A TikiBase consists of pure human-accessible data: Markdown documents stored in
a folder. All you need to use TikiBase is a Markdown or text editor. TikiBase
provides lightweight tooling (this codebase) for automation of peripheral
automation like managing hyperlinks.

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

### best practices

- **avoid deep hierarchies:** Taxonomies are highly debatable, even between your
  current and future self. Keep information organized in a relatively flat
  namespace and retrieve information alphabetically or via fulltext search.

### limitations

- it would be nice to reference other files via **#shortname** rather than
  having to create a Markdown link to the file
