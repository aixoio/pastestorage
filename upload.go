package main

import (
	"github.com/aixoio/pastestorage/converter"
)

func UploadFile(filename string, api_key string) {
	pastes, err := converter.ConvertFileToText(filename)
	if err != nil {
		panic(err)
	}

}
