package parser

import (
	M "http-to-postman/pkg/models"
	"strings"
)

func ParseHttpContent(content string) ([]M.Request, error) {
	var requests []M.Request
	rawRequests := strings.Split(content, "###")

	for _, rawRequest := range rawRequests {
		lines := strings.Split(rawRequest, "\n")
		var headers []M.Header
		if len(lines) < 2 {
			continue
		}
		desc := lines[0]
		method := strings.Split(lines[1], " ")[0]
		url := strings.Split(lines[1], " ")[1]

		if len(lines) >= 3 {
			for _, line := range lines[2:] {
				if strings.Contains(line, "{") {
					break
				}
				if strings.Contains(line, ":") {
					var keyValueArray = strings.Split(line, ":")
					headers = append(headers, M.Header{Key: keyValueArray[0], Value: keyValueArray[1]})
				}

			}
		}
		body := strings.Join(lines[2+len(headers)+1:], "\n")
		requests = append(requests, M.Request{Description: desc, Method: method, Url: url, Headers: headers, Body: body})

	}

	return requests, nil

}
