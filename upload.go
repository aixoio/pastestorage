package main

import (
	"fmt"

	"github.com/aixoio/pastestorage/converter"
)

func UploadFile(filename string) {
	fmt.Println(converter.ConvertFileToText(filename))
}
