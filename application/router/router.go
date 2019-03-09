package router

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "captcha/application/modules/image"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    r.GET("status.html", func(c *gin.Context) {
        c.String(http.StatusOK,"ok\n")
    })

    r.GET("/img/get",image.CaptchaCreate)
    r.GET("/img/check",image.CaptchaVerify)

    return r
}
