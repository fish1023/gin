package tool
import (
    "captcha/application/lib/redisPool"
    "github.com/gomodule/redigo/redis"
)

// Spam 访问频率控制
func Spam(key string,times int,expire int) (b bool){
    value,_ := redis.Int(redisPool.Pool.Deal("incr",key))
    t,_ := redis.Int(redisPool.Pool.Deal("ttl",key))
    if value == 1 || t == -1 {
        redisPool.Pool.Deal("expire",key,expire)
    }

    if value > times {
        b = false
    } else {
        b = true
    }
    return b
}
