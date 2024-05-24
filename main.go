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
	}

}
