package router

import (
    "os"
    "io"
    "net/http"
    "github.com/gin-gonic/gin"
    "captcha/application/modules/image"
    "captcha/application/lib/helper"
)

func init() {
    if e := helper.Env; e == "online" {
        gin.SetMode("release")
        gin.DisableConsoleColor()
        f,_ := os.Create(helper.GetLogPath() + "gin.log")
        gin.DefaultWriter = io.MultiWriter(f)
    }
}

func InitRouter() *gin.Engine {
    r := gin.Default()

    r.GET("status.html", func(c *gin.Context) {
        c.String(http.StatusOK,"ok\n")
    })

    r.GET("/img/get",image.CaptchaCreate)
    r.GET("/img/check",image.CaptchaVerify)

    return r
}
