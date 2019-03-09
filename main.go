package main

import (
    "captcha/application/router"
    "captcha/application/setting"
)

func main() {
    r := router.InitRouter()
    App := setting.Setting.App
    r.Run(":" + App.API.Port)
}
