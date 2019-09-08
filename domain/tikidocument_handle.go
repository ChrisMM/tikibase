package domain

import "strings"

// TikiDocumentHandle is the textual handle of a TikiDocument, i.e. its filename without the '.md' extension.
type TikiDocumentHandle string

// NewHandleFromFileName converts the given filename into its corresponding TikiDocument handle.
func NewHandleFromFileName(filename string) TikiDocumentHandle {
	return TikiDocumentHandle(strings.Replace(filename, ".md", "", 1))
}
