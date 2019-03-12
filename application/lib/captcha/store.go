package captcha

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "captcha/application/lib/redisPool"
    "captcha/application/logger"
    "github.com/mojocn/base64Captcha"
)

type CustomizeRdsStore struct {
    redisClient *redis.Pool
}

var customeStore *CustomizeRdsStore
var l *logger.LogFish

func init() {
    l = &logger.LogFish{Cate:"captcha"}
    customeStore = &CustomizeRdsStore{redisClient:redisPool.RPool}
    base64Captcha.SetCustomStore(customeStore)
}

func (p *CustomizeRdsStore) Set(id string, code string) {
    c := p.redisClient.Get();
    defer c.Close()
    k := CAPTCHA_PREFIX + id;
    errK := ERROR_PREFIX + id;
    _, err := redis.String(c.Do("setex",k,EXPIRE_TIME,code))
    c.Do("setex",errK,EXPIRE_TIME,0)
    if err != nil {
        fmt.Println(err)
        l.Error("set code error id:" + k + " code:" + code + " errmsg:" + err.Error() )
    }
}

func (p *CustomizeRdsStore) Get(id string,clear bool) string {
    c := p.redisClient.Get();
    defer c.Close()
    k := CAPTCHA_PREFIX + id
    code, err := redis.String(c.Do("get",k))
    if err != nil {
        l.Error("get code error id:" + k +  "errmsg:" + err.Error())
    }

    if clear {
		_,err := c.Do("del",k)
		if err != nil {
			l.Error("del code error id:" + id + "errmsg:" + err.Error())
		}
        errK := ERROR_PREFIX + id;
        c.Do("del",errK)
	}
    return code
}

func Del(id string) {
    c := redisPool.RPool.Get();
    defer c.Close()
    k := CAPTCHA_PREFIX + id;
    errK := ERROR_PREFIX + id;
    _,err := c.Do("del",k)
    c.Do("del",errK)
    if err != nil {
        l.Error("del code error id:" + id + "errmsg:" + err.Error())
    }
}
