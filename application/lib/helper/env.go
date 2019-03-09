package helper

import (
	"os"
)

// Env demo online beta
var Env string

// IDC corp bjcc bjyt
var IDC string

func getEnv() string {
	if Env == "" {
		idc := GetIDC()
		if idc == "corp" {
			Env = "demo"
		} else {
			Env = "online"
		}
	}
	return Env
}

// GetIDC idc
func GetIDC() string {
	if IDC == "" {
		IDC = os.Getenv("SYS_IDC_NAME")
		if IDC == "" {
			IDC = "corp"
		}
	}
	return IDC
}
