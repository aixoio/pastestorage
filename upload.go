package main

import (
	"github.com/aixoio/pastestorage/converter"
)

func UploadFile(filename string, api_key string) {
	converter.ConvertFileToText(filename)
}
