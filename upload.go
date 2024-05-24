package main

import (
	"github.com/aixoio/pastestorage/converter"
)

func UploadFile(filename string) {
	converter.ConvertFileToText(filename)
}
