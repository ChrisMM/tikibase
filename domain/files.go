package domain

import (
	"fmt"
)

// DocumentFiles describes the Markdown files in a TikiBase.
type DocumentFiles struct {
	fileNames []string

	// the TikiBase in which these files are located
	tikiBase *TikiBase
}

// Documents provides the documents contained in theses files.
func (df DocumentFiles) Documents() (result Documents, err error) {
	for f := range df.fileNames {
		doc, err := df.tikiBase.LoadDocument(df.fileNames[f])
		if err != nil {
			return result, fmt.Errorf("cannot get documents: %w", err)
		}
		result = append(result, doc)
	}
	return result, nil
}

// FileNames provides the names of the files in this collection.
func (df DocumentFiles) FileNames() []string {
	return df.fileNames
}

// ResourceFiles describes the non-Markdown files in a TikiBase,
// typically images and PDFs.
type ResourceFiles struct {
	fileNames []string
}

// FileNames provides the names of the files in this collection.
func (rf ResourceFiles) FileNames() []string {
	return rf.fileNames
}

// ScaffoldResourceFiles provides resource file collections for testing.
func ScaffoldResourceFiles(files []string) (result ResourceFiles) {
	return ResourceFiles{files}
}
