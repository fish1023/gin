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

type Config struct {
    Id string       `form:"key" json:"id"`
    App string       `form:"app" json:"app"`
    Value string     `form:"code" json:"value"`
    Mode string       `form:"m" json:"m"`
}

func CaptchaCreate(c *gin.Context) {
    var params Config
    appG := app.Gin{C:c}
    c.DefaultQuery("m",captcha.DEFAULT_MODE)
    c.ShouldBind(&params)

    if params.App == "" || params.Id == "" {
        appG.ErrorByCode(retcode.PARAM_ERROR)
        return
    }

    k := params.App + ":" + params.Id
    var config interface{}
    switch params.Mode {
    case captcha.MODE_AUDIO:
        config = captcha.ConfigAudio
    case captcha.MODE_DIGIT:
        config = captcha.ConfigDigit
    case captcha.MODE_CHARACTER:
        config = captcha.ConfigCharacter
    default:
        config = captcha.ConfigDigit
    }

    //GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	_, cap := base64Captcha.GenerateCaptcha(k, config)

	//以base64编码
	base64string := base64Captcha.CaptchaWriteToBase64Encoding(cap)

    appG.Success(base64string)
}

func CaptchaVerify(c *gin.Context) {
    var params Config
    appG := app.Gin{C:c}
    c.ShouldBind(&params)

    if params.App == "" || params.Id == "" {
        appG.ErrorByCode(retcode.PARAM_ERROR)
        return
    }
    k := params.App + ":" +params.Id
    if limited := tool.Spam(captcha.SPAM_PREFIX + k,captcha.MAX_SPAM,captcha.SPAM_TIME);!limited {
        appG.ErrorByCode(retcode.SPAM_ACTION)
        return
    }
    base64Captcha.VerifyCaptcha(k, params.Value)
    verifyResult := base64Captcha.VerifyCaptcha(k, params.Value)
    if !verifyResult {
        if num:= captcha.IncrBadCode(k);num > captcha.ERROR_MAX {
            captcha.Del(k)
        }
    }
    r := strconv.FormatBool(verifyResult)
    appG.Success(r)
}
