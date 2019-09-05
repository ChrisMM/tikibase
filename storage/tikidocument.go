package storage

// TikiDocument represents a MarkDown file in the document base.
type TikiDocument struct {

	// Handle is the unique identifier for this document.
	// It is used to address documents in via links,
	// and it is the filename without directory and extension.
	Handle string

	// the textual content of the document
	Content string
}

// NewTikiDocument creates a new TikiDocument instance in memory.
func NewTikiDocument(handle, content string) TikiDocument {
	return TikiDocument{Handle: handle, Content: content}
}

// GetHandle returns the filename without extension of this TikiDocument.
func (td TikiDocument) GetHandle() string {
	return td.Handle
}
