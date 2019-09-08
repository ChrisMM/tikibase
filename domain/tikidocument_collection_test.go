package domain_test

import (
	"testing"

	"github.com/kevgo/tikibase/domain"
)

func TestTikiDocumentCollectionAdd(t *testing.T) {
	tdc := domain.NewTikiDocumentCollection()
	doc1 := domain.NewTikiDocument("doc1", "")
	doc2 := domain.NewTikiDocument("doc2", "")
	tdc.Add(doc1, doc2)
	documents := tdc.Documents
	if len(documents) != 2 {
		t.Fatalf("wrong number of documents added: %d", len(documents))
	}
}

func TestWikiDocumentCollectionFind(t *testing.T) {
	tdc := domain.NewTikiDocumentCollection()
	doc1 := domain.NewTikiDocument("doc1", "")
	doc2 := domain.NewTikiDocument("doc2", "")
	tdc.Add(doc1, doc2)
	actual, err := tdc.Find("doc2")
	if err != nil {
		t.Fatal(err)
	}
	if actual.Handle() != domain.TikiDocumentHandle("doc2") {
		t.Fatalf("wrong document found: %s", actual.Handle())
	}
}
