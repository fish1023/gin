package captcha

import (
	"captcha/application/lib/redisPool"
	"captcha/application/logger"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/mojocn/base64Captcha"
)

// redis 存储结构
type CustomizeRdsStore struct {
	redisClient *redisPool.PoolManager
}

// redis 存储对象
var customeStore *CustomizeRdsStore
// log 对象
var l *logger.LogFish

func init() {
	l = &logger.LogFish{Cate: "captcha"}
	customeStore = &CustomizeRdsStore{redisClient: redisPool.Pool}
	base64Captcha.SetCustomStore(customeStore)
}

// Set 设置验证码并存储
func (p *CustomizeRdsStore) Set(id string, code string) {
	k := getCaptchaKey(id)
	errK := getErrKey(id)
	_, err := redis.String(p.redisClient.Deal("setex", k, EXPIRE_TIME, code))
	p.redisClient.Deal("setex", errK, EXPIRE_TIME, 0)
	l.Info("set code id:" + k + "code" + code)
	if err != nil {
		fmt.Println(err)
		l.Error("set code error id:" + k + " code:" + code + " errmsg:" + err.Error())
	}
}

// Get 读取验证码&正确后删除
func (p *CustomizeRdsStore) Get(id string, clear bool) string {
	k := getCaptchaKey(id)
	code, err := redis.String(p.redisClient.Deal("get", k))
	if err != nil {
		l.Error("get code error id:" + k + "errmsg:" + err.Error())
	}

	if clear {
		Del(id)
	}
	return code
}

// Del 删除验证码
func Del(id string) {
	k := getCaptchaKey(id)
	errK := getErrKey(id)
	_, err := redisPool.Pool.Deal("del", k)
	redisPool.Pool.Deal("del", errK)
	if err != nil {
		l.Error("del code error id:" + id + "errmsg:" + err.Error())
	}
}
