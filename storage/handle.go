package storage

import "strings"

// Handle is the textual handle of a TikiDocument, i.e. its filename without the '.md' extension.
type Handle string

// NewHandleFromFileName converts the given filename into its corresponding TikiDocument handle.
func NewHandleFromFileName(filename string) Handle {
	return Handle(strings.Replace(filename, ".md", "", 1))
}
