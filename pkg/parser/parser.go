package parser

import (
	"errors"
	"http-to-postman/pkg/models"
	"strings"
)

func ParseHttpContent(content string) ([]models.Request, error) {
	var requests []models.Request
	rawRequests := strings.Split(content, "###")
	for _, rawRequest := range rawRequests {
		if strings.TrimSpace(rawRequest) == "" {
			continue
		}
		lines := sanitizeLines(strings.Split(rawRequest, "\n"))
		desc, method, url := parseDescAndMethodAndUrl(lines)
		headers, body := parseHeadersAndBody(lines[2:])

		requests = append(requests, models.Request{Description: desc, Method: method, Url: url, Headers: headers, Body: body})

	}

	return requests, nil

}

func sanitizeLines(lines []string) []string {
	var sanitized []string
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			sanitized = append(sanitized, trimmedLine)
		}
	}
	return sanitized
}

func parseDescAndMethodAndUrl(lines []string) (string, string, string) {
	if len(lines) < 2 {
		panic(errors.New("must have at least two lines"))
	}
	desc := lines[0]
	methodAndUrl := strings.Split(lines[1], " ")

	return desc, methodAndUrl[0], methodAndUrl[1]
}

func parseHeadersAndBody(lines []string) ([]models.Header, string) {
	var headers []models.Header
	for _, line := range lines {
		if strings.Contains(line, "{") {
			break
		}
		if strings.Contains(line, ":") {
			var keyValueArray = strings.Split(line, ":")
			TrimmedKey := strings.TrimSpace(keyValueArray[0])
			TrimmedValue := strings.TrimSpace(keyValueArray[1])
			headers = append(headers, models.Header{Key: TrimmedKey, Value: TrimmedValue})
		}

	}

	var bodyLines []string

	for _, line := range lines[len(headers):] {
		if strings.Contains(line, "}") {
			bodyLines = append(bodyLines, line)
			break
		}
		bodyLines = append(bodyLines, line)

	}

	body := strings.Join(bodyLines, "\n")

	return headers, body
}
