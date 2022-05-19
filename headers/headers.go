package headers

import "net/http"

func ParseHeaders(headers http.Header) string {
	var output string
	for key, value := range headers {
		output += key + ": "
		for _, item := range value {
			output += item
		}
		output += "\n"
	}
	return output
}
