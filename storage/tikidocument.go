package storage

import "strings"

// TikiDocument represents a MarkDown file in the document base.
type TikiDocument struct {

	// handle is the unique identifier for this document.
	// It is used to address documents in via links,
	// and it is the filename without directory and extension.
	handle Handle

	// the textual content of the document
	content string
}

// AllSections returns all the TikiSections that make up this document,
// including the title section.
func (td TikiDocument) AllSections() []TikiSection {
	result := []TikiSection{}
	var tsb TikiSectionBuilder
	for i, line := range strings.Split(td.content, "\n") {
		if strings.HasPrefix(line, "#") {
			if i > 0 {
				result = append(result, tsb.Section())
			}
			tsb = NewTikiSectionBuilder(line)
		} else {
			tsb.AddLine(line)
		}
	}
	result = append(result, tsb.Section())
	return result
}

// FilePath returns the file path (handle + extension) of this TikiDocument.
func (td TikiDocument) FilePath() string {
	return string(td.handle) + ".md"
}

// Handle returns the filename without extension of this TikiDocument.
func (td TikiDocument) Handle() Handle {
	return td.handle
}

// NewTikiDocument creates a new TikiDocument instance in memory.
func NewTikiDocument(handle Handle, content string) TikiDocument {
	return TikiDocument{handle: handle, content: content}
}
