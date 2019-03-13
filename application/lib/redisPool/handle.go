package redisPool

import(
	"github.com/gomodule/redigo/redis"
)

var CMD_MASTER_MAP = map[string]int{
	"append":1, "decr":1, "decrby":1, "getset":1, "incr":1, "incrby":1, "incrbyfloat":1,
        "mset":1, "msetnx":1, "set":1, "setbit":1, "setex":1, "psetex":1, "setnx":1, "setrange":1, "exists":1,
        "del":1, "delete":1, "expire":1, "settimeout":1, "pexpire":1, "expireat":1, "pexpireat":1,
        "migrate":1, "move":1, "persist":1, "rename":1, "renamekey":1, "renamenx":1, "restore":1,
        "hdel":1, "hincrby":1, "hset":1, "hsetnx":1, "hmset":1,
        "blpop":1, "brpop":1, "brpoplpush":1, "linsert":1, "lpop":1, "lpush":1, "lpushx":1, "lrem":1,
        "lremove":1, "lset":1, "ltrim":1, "listtrim":1, "rpop":1, "rpoplpush":1, "rpush":1, "rpushx":1,
        "sadd":1, "sdiffstore":1, "sinterstore":1, "smove":1, "spop":1, "srem":1, "sremove":1, "sunionstore":1,
        "zadd":1, "zincrby":1, "zinter":1, "zrem":1, "zdelete":1, "zremrangebyrank":1,
        "zdeleterangebyrank":1, "zremrangebyscore":1, "zdeleterangebyscore":1, "zunion":1,
}

// Deal 对外接口 对Do进行包裹 自动选择连接池
func (pool *PoolManager)Deal(method string,params ... interface{}) (reply interface{},err error){
	p := pool.get(method)
	c := p.Get()
	defer c.Close()
	reply,err = c.Do(method,params ...)
	return
}

// 选择连接池
func (pool *PoolManager)get(method string) (p *redis.Pool){
	if CMD_MASTER_MAP[method] == 1 {
		p = pool.master
	} else {
		p = pool.slave
	}
	return
}
