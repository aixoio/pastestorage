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
		if len(os.Args) < 7 {
			fmt.Println("Usage: pastestorage upload <FILENAME> <API_KEY> <USERNAME> <PASSWORD> <AES_KEY>")
			os.Exit(1)
		}
		UploadFile(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		return
	}

	if strings.Compare(strings.ToLower(os.Args[1]), "download") == 0 {
		if len(os.Args) < 7 {
			fmt.Println("Usage: pastestorage download <API_KEY> <USERNAME> <PASSWORD> <AES_KEY> <LINK>")
			os.Exit(1)
		}
		DownloadFile(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		return
	}

}
