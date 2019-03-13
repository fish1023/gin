package captcha
import (
    "captcha/application/lib/redisPool"
    "github.com/gomodule/redigo/redis"
)

// 累计错误次数增长
func IncrBadCode(id string) int{
    k := ERROR_PREFIX + id
    num,_ := redis.Int(redisPool.Pool.Deal("incr",k))
    return num
}
