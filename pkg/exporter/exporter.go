package exporter

import (
	"encoding/json"
	"fmt"
	"http-to-postman/pkg/models"
	"os"
)

func ExportPostmanCollections(collections models.PostmanCollection, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(collections)
	if err != nil {
		return fmt.Errorf("failed to encode collections: %v", err)
	}

	fmt.Printf("Successfully wrote collection to %s\n", filename)
	return nil
}
