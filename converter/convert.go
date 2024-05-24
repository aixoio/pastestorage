package converter

import (
	"encoding/base64"
	"math"
	"os"
	"sync"
)

const KB_PER_SPLIT_IN_BYTES = 512000

func ConvertFileToText(filename string) ([]string, error) {
	file_dat, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	unsplited := base64.StdEncoding.EncodeToString(file_dat)

	split_count := int(math.Ceil(float64(len(unsplited)) / float64(KB_PER_SPLIT_IN_BYTES)))
	out := make([]string, split_count)
	var wg sync.WaitGroup

	wg.Add(split_count)

	for i := 0; i < split_count; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * KB_PER_SPLIT_IN_BYTES
			end := start + KB_PER_SPLIT_IN_BYTES
			if end > len(unsplited) {
				end = len(unsplited)
			}
			out[i] = unsplited[start:end]
		}(i)
	}

	wg.Wait()

	return out, nil
}
