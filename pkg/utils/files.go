package utils

import (
	"os"
)

func ReadHttpFile(src string) (string, error) {
	data, err := os.ReadFile(src)

	if err != nil {
		panic(err)
	}

	return string(data), err
}
