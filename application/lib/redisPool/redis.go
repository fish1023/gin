package redisPool

import (
	"captcha/application/lib/helper"
	"captcha/application/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

// PoolManager 连接池管理 包含读写 读两个池子
type PoolManager struct {
	master *redis.Pool
	slave  *redis.Pool
}

// Pool 对外开放的对象
var Pool = &PoolManager{}

// 初始化连接池
func init() {
	conf := setting.Setting.Redis
	idc := helper.IDC
	idcConf := conf[idc]
	Pool.master = newPool(&idcConf.Master)
	Pool.slave = newPool(&idcConf.Slave)
}

// 构建连接池
func newPool(s *setting.RedisServer) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Host+":"+s.Port)
			if err != nil {
				return nil, err
			}

			c.Do("AUTH", s.Pass)
			return c, nil
		},
	}
	return pool
}
