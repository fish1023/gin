package helper
import (
    "captcha/application/setting"
)

var logPath string

var port string

var redis interface{}

func GetPort() string {
    p := "80"

    if port != "" {
        p = port
    } else {
        c := setting.Setting.App[Env]
        if c.Port != "" {
            p = c.Port
        }
    }

    return p
}

func GetLogPath() string {
    lp := "/data/log/"

    if logPath != "" {
        lp = logPath
    } else {
        c := setting.Setting.App[Env]
        if c.LogPath != "" {
            lp = c.LogPath
        }
    }
    return lp
}
