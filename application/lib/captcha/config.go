package captcha
import (
    "github.com/mojocn/base64Captcha"
)

const MODE_DIGIT string = "0"
const MODE_CHARACTER string = "1"
const MODE_AUDIO string = "2"

const EXPIRE_TIME = 300
const DEFAULT_MODE string = "1"
const DEFAULT_LEN int = 4
const MAX_SPAM = 5
const SPAM_TIME = 5
const CAPTCHA_PREFIX = "captcha:"
const SPAM_PREFIX = "captcha:spam:"
const ERROR_PREFIX = "captcha:err:"
const ERROR_MAX = 10

var configAudio = base64Captcha.ConfigAudio{
    CaptchaLen: DEFAULT_LEN,
    Language:   "zh",
}

var configCharacter = base64Captcha.ConfigCharacter{
    Height:             60,
    Width:              240,
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

var configDigit = base64Captcha.ConfigDigit{
    Height:     80,
    Width:      240,
    MaxSkew:    0.7,
    DotCount:   80,
    CaptchaLen: DEFAULT_LEN,
}

func GetConfigByMode(mode string) (config interface{}) {
    switch mode {
    case MODE_AUDIO:
        config = configAudio
    case MODE_DIGIT:
        config = configDigit
    case MODE_CHARACTER:
        config = configCharacter
    default:
        config = configDigit
    }
    return
}
