package parser

import (
	"fmt"
	M "http-to-postman/pkg/models"
	"strings"
)

func ParseHttpContent(content string) ([]M.Request, error) {
	var requests []M.Request
	rawRequests := strings.Split(content, "###")

	fmt.Print(rawRequests)

	return requests, nil

}
