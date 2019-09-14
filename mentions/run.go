package mentions

import (
	"fmt"
	"log"

	"github.com/kevgo/tikibase/domain"
	"github.com/pkg/errors"
)

// Run executes the "mentions" command in the given directory.
// Channel-based data flow processing setup:
// Goroutine 1: reads all Documents and sends out one filename at a time via  the "TikiDocs" channel.
//              when it is done, it closes the channel.
// Workerpool 2: reads one Document filename at a time from the "TikiDocs" channel
//               and finds all the TikiLinks in that document.
//               It sends each link out via the "TikiLinks" channel.
//               When it sees the "TikiDocs" channel closing, it closes the "TikiLinks" channel.
// Goroutine 3: reads one TikiLink at a time from the "TikiLinks" channel
//              and builds up a map of documents to incoming links.
//              When the "TikiLinks" channel closes, it sends out the map via the "IncomingLinks" channel.
// Goroutine 4: reads the map from the "IncomingLinks" channel.
//              Once it has that, it reads one TikiDoc path at a time from a copy of the "TikiDocs" channel
//              and sends an "UpdateTikiDoc" command containing a Document path and list of incomings links to a channel.
// Workerpool 5: receives one "UpdateTikiDoc" command at a time, compiles the "mentions" section,
//               and sends that out as a "UpdateMentionsSectionInDoc" command via a channel
// Workerpool 6: receives one "UpdateMentionsSectionInDoc" command at a time, updates the file on disk,
//               and sends out a "TikiDocUpdated" message via a channel
// Goroutine 7: collects all "TikiDocUpdated" messages, prints the result, and ends the program.
//
//
// This can be built over several steps:
// Step 1: serial version
// Step 2: single-threaded workers connected by channels
// Step 3: some steps are implemented as worker pools.
//         We need to build a "WorkerPool" struct that has an incoming channel
//         and an outgoing channel. It spins up workers that read from the incoming channel and send results to the outgoing channel.
//         When its input channel closes, the worker simply stop and end their a WaitGroup.
//         When the waitgroup is done, the WorkerPool closes the outgoing channel.
//
// This feels too low-level, and complex for timeless code.
//
//
// A higher-level abstraction might be a MapReduce setup:
// Step 1: expand a directory name to a list of filenames
// Step 2: in parallel, map over all filenames to get their TikiLinks.
// Step 3: reduce all found TikiLinks to the LinkMap of Documentsn and incoming TikiLinks.
//
// Step 4: expand the directory name to a list of filenames again
// Step 5: in parallel, map over all filenames, render the "mentions" section, and update the file
// Step 6: reduce to the number of files processed and quit
//
// The map steps would be executed in parallel,	using a setup like http://www.golangpatterns.info/concurrency/parallel-for-loop
// but with a WaitGroup instead of the channel-based semaphore.
//
//
// The most integrated version:
// Make `tb.TikiLinks()` work in parallel internally by using a parallel range construct
// that we have to build for this based on http://www.golangpatterns.info/concurrency/parallel-for-loop.
//
// Or, assuming that channels are high-level, first-grade language constructs
// to encourage concurrent programming,
// have `tb.TikiLinks()` return a channel of TikiLinks rather than a slice.
// The problem is that channels are one-time use, though.
// They seem for communication, not for data storage.
// I guess if one has to store results, one has to save them from the channel into a list
// using some form of channel to list reader.
//
// For timelessness it's probably better to keep it as simple as possible, i.e. single-threaded.
func Run(dir string) error {
	tb, err := domain.NewTikiBase(dir)
	if err != nil {
		return err
	}

	docs, err := tb.Documents()
	if err != nil {
		return errors.Wrapf(err, "cannot get documents of TikiBase")
	}
	fmt.Printf("%d documents found\n", len(docs))

	allLinks, err := docs.TikiLinks()
	if err != nil {
		return errors.Wrapf(err, "cannot get links of TikiBase")
	}
	fmt.Printf("%d total links found\n", len(allLinks))

	linksToDocs := allLinks.GroupByTarget()
	fmt.Printf("%d linked documents found\n", len(linksToDocs))

	for _, doc := range docs {
		doc := doc // need to pin the loop variable so that we can use it in functions
		fileName := doc.FileName()
		linksToDoc := linksToDocs[fileName]
		fmt.Printf("- %s: %d references\n", fileName, len(linksToDoc))
		mentionsSection := RenderMentionsSection(linksToDoc, &doc)
		doc2 := doc.AppendSection(mentionsSection)
		err := tb.SaveDocument(doc2)
		if err != nil {
			log.Fatalf("cannot update document %s: %v", fileName, err)
		}
	}

	return nil
}
