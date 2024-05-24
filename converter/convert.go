package converter

import (
	"encoding/base64"
	"math"
	"os"
)

const KB_PER_SPLIT_IN_BYTES = 512000

func ConvertFileToText(filename string) ([]string, error) {
	file_dat, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	unsplited := base64.StdEncoding.EncodeToString(file_dat)

	out := []string{}

	for i := 0; i < int(math.Ceil(float64(len(unsplited))/float64(KB_PER_SPLIT_IN_BYTES))); i++ {
		str := ""

		start := i * KB_PER_SPLIT_IN_BYTES
		end := start + KB_PER_SPLIT_IN_BYTES
		if end > len(unsplited) {
			end = len(unsplited)
		}

		for j := start; j < end; j++ {
			str += string(unsplited[j])
		}
		out = append(out, str)
	}

	return out, nil
}
