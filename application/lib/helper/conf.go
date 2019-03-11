package helper

import (
	"os"
)

// Env demo online beta
var Env string

// IDC corp bjcc bjyt
var IDC string

func init() {
    IDC = os.Getenv("SYS_IDC_NAME")
    if IDC == "" {
        IDC = "corp"
    }

    if IDC == "corp" {
        Env = "dev"
    } else {
        Env = "online"
    }
}
