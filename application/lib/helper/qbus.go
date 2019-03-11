package helper

import (
    "os"
    "fmt"
)

func GetQBusConf() string {
    fmt.Println("IDC is :" + IDC)
	iniF := "conf/bjyt_pub_infra.ini"
	_, err := os.Stat("conf/" + IDC + "_pub_infra.ini")
	if err == nil {
		iniF = "conf/" + IDC + "_pub_infra.ini"
	}
	return iniF
}
