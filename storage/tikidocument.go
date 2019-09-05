package storage

// TikiDocument represents a MarkDown file in the document base.
type TikiDocument struct {

	// handle is the unique identifier for this document.
	// It is used to address documents in via links,
	// and it is the filename without directory and extension.
	handle string

	// the textual content of the document
	content string
}

// NewTikiDocument creates a new TikiDocument instance in memory.
func NewTikiDocument(handle, content string) TikiDocument {
	return TikiDocument{handle: handle, content: content}
}

// Handle returns the filename without extension of this TikiDocument.
func (td TikiDocument) Handle() string {
	return td.handle
}
