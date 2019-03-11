package main

import (
    "captcha/application/router"
    "captcha/application/lib/helper"
)

func main() {
    r := router.InitRouter()
    port := helper.GetPort()
    r.Run(":" + port)
}
