package storage_test

import (
	"testing"

	"github.com/kevgo/tikibase/storage"
)

func TestTikiDocumentCollectionAdd(t *testing.T) {
	tdc := storage.NewTikiDocumentCollection()
	doc1 := storage.NewTikiDocument("doc1", "")
	doc2 := storage.NewTikiDocument("doc2", "")
	tdc.Add(doc1, doc2)
	documents := tdc.Documents
	if len(documents) != 2 {
		t.Fatalf("wrong number of documents added: %d", len(documents))
	}
}

func TestWikiDocumentCollectionFind(t *testing.T) {
	tdc := storage.NewTikiDocumentCollection()
	doc1 := storage.NewTikiDocument("doc1", "")
	doc2 := storage.NewTikiDocument("doc2", "")
	tdc.Add(doc1, doc2)
	actual, err := tdc.Find("doc2")
	if err != nil {
		t.Fatal(err)
	}
	if !actual.HasHandle("doc2") {
		t.Fatalf("wrong document found: %s", actual.Handle())
	}
}
