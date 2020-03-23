package linkify

import "github.com/kevgo/tikibase/domain"

// DocLinks provides all names and targets for the given documents.
func DocLinks(docs domain.Documents) (result map[string]string, err error) {
	result = make(map[string]string) // text -> filename
	for d := range docs {
		names, err := docs[d].Names()
		if err != nil {
			return result, err
		}
		for n := range names {
			result[names[n]] = docs[d].FileName()
		}
	}
	return result, nil
}
