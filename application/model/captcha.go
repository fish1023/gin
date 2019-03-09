package model

import (
    "captcha/application/lib/redisPool"
    "captcha/application/lib/logger"
    "captcha/application/lib/tool"
    "github.com/gomodule/redigo/redis"
)

func (p *tool.CustomizeRdsStore) Set(id string, code string) {
    c := p.redisClient.Get();
    defer c.Close()
    _,err := c.DO('set',id,code)
    if(err != nil) {
        logger.Error("set code error id:" + id + " code:" + code + " errmsg:" + err)
    }
}

func (p *tool.CustomizeRdsStore) Get(id string,clear bool) (value string) {
    c := p.redisClient.Get();
    defer c.Close()
    code,err := redis.String(c.DO('get',id))
    if(err != nil) {
        logger.Error("get code error id:" + id +  "errmsg:" + err)
    }

    if clear {
		err := c.Do('del',id)
		if err != nil {
			panic(err)
		}
	}
    return code
}
