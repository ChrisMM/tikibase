package linkify

import (
	"fmt"

	"github.com/jinzhu/inflection"
	"github.com/kevgo/tikibase/domain"
	"github.com/kevgo/tikibase/helpers"
)

type docMapping struct {
	lookFor     string
	replaceWith string
	file        string
}

func docsMappings(docs domain.Documents) (result []docMapping, err error) {
	docsNames := make(map[string]string) // doc name -> filename
	for d := range docs {
		docNames, err := docs[d].Names()
		if err != nil {
			return result, err
		}
		for dn := range docNames {
			docsNames[docNames[dn]] = docs[d].FileName()
		}
	}
	keys := make([]string, 0, len(docsNames))
	for docName := range docsNames {
		keys = append(keys, docName)
	}
	helpers.LongestFirst(keys)
	result = make([]docMapping, 0, len(keys))
	for k := range keys {
		result = append(result, docMapping{
			lookFor:     keys[k],
			replaceWith: fmt.Sprintf("[%s](%s)", keys[k], docsNames[keys[k]]),
			file:        docsNames[keys[k]],
		})
		plural := inflection.Plural(keys[k])
		if plural != keys[k] {
			result = append(result, docMapping{
				lookFor:     plural,
				replaceWith: fmt.Sprintf("[%s](%s)", plural, docsNames[keys[k]]),
				file:        docsNames[keys[k]],
			})
		}
	}
	return result, nil
}
