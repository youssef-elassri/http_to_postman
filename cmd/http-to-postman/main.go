package main

import (
	P "http-to-postman/pkg/parser"
	U "http-to-postman/pkg/utils"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("Please provide the .http file, Use './http-to-postman test.http'.")
	//	return
	//}

	// The first argument after the program name (os.Args[0]) is the command
	// httpFile := os.Args[1]
	httpFile := "/Users/imediadevice/Desktop/open-source/http-to-postman/cmd/http-to-postman/test.http"

	content, _ := U.ReadHttpFile(httpFile)
	_, err := P.ParseHttpContent(content)
	if err != nil {
		return
	}

}
