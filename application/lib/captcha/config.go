package captcha
import (
    "strconv"
    "github.com/mojocn/base64Captcha"
)

const MODE_DIGIT string = "0"
const MODE_CHARACTER string = "1"
const MODE_AUDIO string = "2"

const EXPIRE_TIME = 300
const DEFAULT_MODE string = "1"
const DEFAULT_LEN int = 4
const DEFAULT_WIDTH string = "172"
const DEFAULT_HEIGHT string = "60"

const MAX_SPAM = 5
const SPAM_TIME = 5
const CAPTCHA_PREFIX = "captcha:"
const SPAM_PREFIX = "captcha:spam:"
const ERROR_PREFIX = "captcha:err:"
const ERROR_MAX = 10


type Config struct {
    Id string       `form:"key" json:"id"`
    App string       `form:"app" json:"app"`
    Value string     `form:"code" json:"value"`
    Width string      `form:"width" json:"width"`
    Height string     `form:"height" json:"height"`
    Len  string       `form:"length" json:"length"`
    Mode string       `form:"m" json:"m"`
}

var configAudio = base64Captcha.ConfigAudio{
    CaptchaLen: DEFAULT_LEN,
    Language:   "zh",
}

func getConfigCharacter(width int,height int) interface{}{
    return base64Captcha.ConfigCharacter{
        Height:             height,
        Width:              width,
        //const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
        Mode:               base64Captcha.CaptchaModeNumberAlphabet,
        ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
        ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
        IsUseSimpleFont:    true,
        IsShowHollowLine:   false,
        IsShowNoiseDot:     false,
        IsShowNoiseText:    false,
        IsShowSlimeLine:    false,
        IsShowSineLine:     false,
        CaptchaLen:         DEFAULT_LEN,
    }

}

func getConfigDigit(width int,height int) interface{}{
    return base64Captcha.ConfigDigit{
        Height:     height,
        Width:      width,
        MaxSkew:    0.7,
        DotCount:   80,
        CaptchaLen: DEFAULT_LEN,
    }

}

func GetConfig(params *Config) (config interface{}) {
    width := params.Width
    height := params.Height
    if width == "" {
        width = DEFAULT_WIDTH
    }
    if height == "" {
        height = DEFAULT_HEIGHT
    }
    w,_ := strconv.Atoi(width)
    h,_ := strconv.Atoi(height)

    switch params.Mode {
    case MODE_AUDIO:
        config = configAudio
    case MODE_DIGIT:
        config = getConfigDigit(w,h)
    case MODE_CHARACTER:
        config = getConfigCharacter(w,h)
    default:
        config =  getConfigDigit(w,h)
    }
    return
}
