package headers

import (
	"fmt"
	"net/http"
	"sort"
)

type Header struct {
	Name   string
	Values []string
}

func ParseHeaders(headers http.Header) string {
	var output string
	var _headers []Header

	//flatten r.Headers map to Header struct slice
	for key, value := range headers {
		newHeader := Header{
			Name:   key,
			Values: value,
		}
		_headers = append(_headers, newHeader)
	}

	//sort Header slice by Name
	sort.Slice(_headers, func(i, j int) bool {
		return _headers[i].Name < _headers[j].Name
	})

	for _, value := range _headers {
		output += fmt.Sprintf("%s: ", value.Name)
		for _, field := range value.Values {
			output += fmt.Sprintf("%s ", field)
		}
		output += "\n"
	}

	return output
}
