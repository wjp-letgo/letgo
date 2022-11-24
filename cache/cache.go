package cache

import (
	"github.com/wjpxxx/letgo/cache/filecache"
	"github.com/wjpxxx/letgo/cache/icache"
	"github.com/wjpxxx/letgo/cache/memcache"
	"github.com/wjpxxx/letgo/cache/redis"
	"sync"
)

var cacheList map[string]icache.ICacher

var queueList map[string]icache.Quequer

var rds redis.Rediser

var rdsLock sync.Mutex

//Register 注册缓存对象
func Register(name string, cacher icache.ICacher) {
	if cacheList==nil{
		cacheList=make(map[string]icache.ICacher)
	}
	cacheList[name]=cacher
}
//NewCache 新建一个缓存对象
func NewCache(name string) icache.ICacher{
	return cacheList[name]
}
//NewRedis 新建一个redis对象
func NewRedis()redis.Rediser{
	rdsLock.Lock()
	defer rdsLock.Unlock()
	if rds!=nil{
		return rds
	}else{
		rds=redis.NewRedis()
		return rds
	}
	
}
//GetRedis
func GetRedis()redis.Master{
	return NewRedis().Master()
}
//GetSlaveRedis
func GetSlaveRedis(name string)redis.Master{
	return NewRedis().SlaveByName(name)
}

//NewQueue
func NewQueue(name string) icache.Quequer{
	return queueList[name]
}
//Register 注册缓存对象
func RegisterQueue(name string, cacher icache.Quequer) {
	if queueList==nil{
		queueList=make(map[string]icache.Quequer)
	}
	queueList[name]=cacher
}
//init 注册系统的
func init(){
	rds=redis.NewRedis()
	Register("redis", rds.Master())
	Register("file", filecache.NewFileCache())
	Register("memcache", memcache.NewMemcache())
	RegisterQueue("redis", rds.Master())
}