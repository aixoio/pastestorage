package converter

import (
	"encoding/base64"
	"math"
	"os"
	"sync"

	"github.com/aixoio/pastestorage/aes"
	"github.com/aixoio/pastestorage/hashing"
)

const KB_PER_SPLIT_IN_BYTES = 512000

func ConvertFileToText(filename, aes_key string) ([]string, error) {
	file_dat, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	key := hashing.Sha256_to_bytes([]byte(aes_key))
	dat, _ := aes.AesGCMEncrypt(key, file_dat)

	unsplited := base64.StdEncoding.EncodeToString(dat)

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

func ConvertTextToFile(filename, aes_key string, text []string) {
	unsplited := ""

	for i := 0; i < len(text); i++ {
		unsplited += text[i]
	}

	dat, _ := base64.StdEncoding.DecodeString(unsplited)
	key := hashing.Sha256_to_bytes([]byte(aes_key))
	real_dat, _ := aes.AesGCMDecrypt(key, dat)
	os.WriteFile(filename, real_dat, 0644)
}
