package helper

import (
	"captcha/application/logger"
	"os"
)

func GetQBusConf() string {
	logger := &logger.LogFish{Cate: "qbus"}
	IDC := GetIDC()
	logger.Debug("IDC is :" + IDC)
	iniF := "conf/bjyt_pub_infra.ini"
	_, err := os.Stat("conf/" + IDC + "_pub_infra.ini")
	if err == nil {
		iniF = "conf/" + IDC + "_pub_infra.ini"
	}
	return iniF
}
