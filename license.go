package main

import (
	_ "embed"
	"fmt"
)

//go:embed LICENSE
var full_license string

func ShowLicenseMessage() {
	fmt.Println("This program is under the AGPL-3.0 license use pastestorage license to find out more")
}

func ShowFullLicense() {
	ShowLicenseMessage()
	fmt.Println(full_license)
}
