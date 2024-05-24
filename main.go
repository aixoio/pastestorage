package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	ShowLicenseMessage()
	if len(os.Args) < 2 {
		fmt.Println("Usage: pastestorage <MODE>")
		os.Exit(1)
	}

	if strings.Compare(strings.ToLower(os.Args[1]), "license") == 0 {
		ShowFullLicense()
		return
	}

	if strings.Compare(strings.ToLower(os.Args[1]), "upload") == 0 {
		if len(os.Args) < 3 {
			fmt.Println("Usage: pastestorage upload <FILENAME>")
			os.Exit(1)
		}
		UploadFile(os.Args[2])
		return
	}

}
