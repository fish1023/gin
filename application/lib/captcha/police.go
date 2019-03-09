package captcha
import (
    "captcha/application/lib/redisPool"
    "github.com/gomodule/redigo/redis"
)

func IncrBadCode(id string) int{
    c := redisPool.RPool.Get();
    defer c.Close()
    k := ERROR_PREFIX + id
    num,_ := redis.Int(c.Do("incr",k))
    return num
}
