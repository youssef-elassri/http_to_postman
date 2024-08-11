package main

import (
	"flag"
	"fmt"
	"http-to-postman/pkg/converter"
	"http-to-postman/pkg/exporter"
	"http-to-postman/pkg/parser"
	"http-to-postman/pkg/utils"
	"os"
)

func main() {

	httpFile := flag.String("http", "", "Path to the HTTP file")
	outputFile := flag.String("output", "postman_collection.json", "Path to the output file")
	help := flag.Bool("help", false, "Display help message")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if *httpFile == "" {
		fmt.Println("Error: HTTP file path must be provided.")
		flag.Usage()
		os.Exit(1)
	}

	content, _ := utils.ReadHttpFile(*httpFile)
	requests, err := parser.ParseHttpContent(content)
	if err != nil {
		return
	}

	collection := converter.ConvertToPostmanCollection(requests)

	// outputFile := strings.Split(httpFile, ".")[0] + ".json"

	err = exporter.ExportPostmanCollections(collection, *outputFile)
	if err != nil {
		fmt.Printf("Error exporting collection: %v\n", err)
		os.Exit(1)
	}
}
