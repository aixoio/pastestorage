package converter

import (
	"encoding/base64"
	"os"
)

const KB_PER_SPLIT_IN_BYTES = 512000

func ConvertFileToText(filename string) ([]string, error) {
	file_dat, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	unsplited := base64.StdEncoding.EncodeToString(file_dat)

	return []string{}, nil
}
