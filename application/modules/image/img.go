package image

import (
    //"fmt"
    "strconv"
    "github.com/mojocn/base64Captcha"
    "github.com/gin-gonic/gin"
    "captcha/application/lib/tool"
    "captcha/application/lib/app"
    "captcha/application/lib/retcode"
    "captcha/application/lib/captcha"
)


func CaptchaCreate(c *gin.Context) {
    var params captcha.Config
    appG := app.Gin{C:c}
    c.DefaultQuery("m",captcha.DEFAULT_MODE)
    c.ShouldBind(&params)

    if params.App == "" || params.Id == "" {
        appG.ErrorByCode(retcode.PARAM_ERROR)
        return
    }

    k := params.App + ":" + params.Id
    config := captcha.GetConfig(&params)

    //GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	_, cap := base64Captcha.GenerateCaptcha(k, config)

	//以base64编码
	base64string := base64Captcha.CaptchaWriteToBase64Encoding(cap)

    appG.Success(base64string)
}

func CaptchaVerify(c *gin.Context) {
    var params captcha.Config
    appG := app.Gin{C:c}
    c.ShouldBind(&params)

    if params.App == "" || params.Id == "" {
        appG.ErrorByCode(retcode.PARAM_ERROR)
        return
    }
    k := params.App + ":" +params.Id
    // 频率限制
    if limited := tool.Spam(captcha.SPAM_PREFIX + k,captcha.MAX_SPAM,captcha.SPAM_TIME);!limited {
        appG.ErrorByCode(retcode.SPAM_ACTION)
        return
    }

    verifyResult := base64Captcha.VerifyCaptcha(k, params.Value)
    if !verifyResult {
        // 错误次数限制
        if num:= captcha.IncrBadCode(k);num > captcha.ERROR_MAX {
            captcha.Del(k)
        }
    }
    r := strconv.FormatBool(verifyResult)
    appG.Success(r)
}
