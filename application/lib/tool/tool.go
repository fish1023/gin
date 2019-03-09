package tool
import (
    "captcha/application/lib/redisPool"
    "github.com/gomodule/redigo/redis"
)

func Spam(key string,times int,expire int) (b bool){
    c := redisPool.RPool.Get()
    defer c.Close()
    value,_ := redis.Int(c.Do("incr",key))
    t,_ := redis.Int(c.Do("ttl",key))
    if value == 1 || t == -1 {
        c.Do("expire",key,expire)
    }

    if value > times {
        b = false
    } else {
        b = true
    }
    return b
}
