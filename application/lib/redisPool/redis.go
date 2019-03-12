package redisPool

import (
	"time"
    "fmt"
	"captcha/application/lib/helper"
	"captcha/application/setting"

	"github.com/gomodule/redigo/redis"
)

// RPool redis pool
var RPool *redis.Pool

// Setup 处理redis命令
func init() {
	conf := setting.Setting.Redis
	idc := helper.IDC
	idcConf := conf[idc]
    fmt.Printf("redis config %+v",idcConf)
	// RPool = &redis.Pool{
	// 	MaxIdle:     1,
	// 	MaxActive:   10,
	// 	IdleTimeout: 180 * time.Second,
	// 	Dial: func() (redis.Conn, error) {
	// 		c, err := redis.Dial("tcp", rConf.Host+":"+rConf.Port)
	// 		if err != nil {
	// 			return nil, err
	// 		}

	// 		c.Do("AUTH", rConf.Pass)
    //         return c, nil
	// 	},
	// }
    fmt.Println("redis is ready")
}
