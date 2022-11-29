package memcache

import (
	"fmt"
	"sync"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/wjp-letgo/letgo/cache/icache"
	"github.com/wjp-letgo/letgo/file"
	"github.com/wjp-letgo/letgo/lib"
)

//全局实现者
var pool MemcachePooler
var poolLock sync.Mutex

//Memcacher 操作接口 
type Memcacher interface {
	Master()icache.ICacher
	Slave()icache.ICacher
	SlaveByName(name string)icache.ICacher
}

//Memcach Memcach对象
type Memcach struct {
	pool MemcachePooler
	isMaster bool
	slaveName string

}
//Master 主redis
func (r *Memcach)Master()icache.ICacher{
	r.isMaster=true
	return r
}
//Slave 从redis
func (r *Memcach)Slave()icache.ICacher{
	r.isMaster=false
	r.slaveName=""
	return r
}
//SlaveByName 从redis 通过名称
func (r *Memcach)SlaveByName(name string)icache.ICacher{
	r.isMaster=false
	r.slaveName=name
	return r
}
//getMemcache 获得memcache
func (r *Memcach)getMemcache() *memcache.Client{
	if r.isMaster{
		return r.pool.GetMaster()
	}else{
		if r.slaveName=="" {
			return r.pool.GetSlave()
		}else{
			return r.pool.GetSlaveByName(r.slaveName)
		}
		
	}
}
//Set
func (r *Memcach)Set(key string, value interface{}, overtime int64) bool{
	err:=r.getMemcache().Set(&memcache.Item{
		Key:key,
		Value:lib.Serialize(value),
		Expiration:int32(overtime),
	})
	if err!=nil{
		return false
	}
	return true
}
//Get
func (r *Memcach)Get(key string, value interface{}) bool{
	item,err:=r.getMemcache().Get(key)
	if err!=nil{
		return false
	}
	lib.UnSerialize(item.Value,value)
	return true
}
//Del
func (r *Memcach)Del(key string) bool{
	err:=r.getMemcache().Delete(key)
	if err!=nil{
		return false
	}
	return true
}
//FlushDB
func (r *Memcach)FlushDB() bool{
	err:=r.getMemcache().FlushAll()
	if err!=nil{
		return false
	}
	return true
}

//SetPool 设置池
func (r *Memcach)SetPool(pooler MemcachePooler)icache.ICacher {
	r.pool=pooler
	return r
}

//NewRedis 新建一个redis
func NewMemcache()icache.ICacher{
	memcacheFile:="config/memcached.config"
	cfgFile:=file.GetContent(memcacheFile)
	var config MemCacheConnect
	if cfgFile==""{
		var slaves []SlaveDB=make([]SlaveDB, 1)
		master:=SlaveDB{
			Name:"name",
			Host:"127.0.0.1",
			Port:"11211",
			MaxIdle:20,
			IdleTimeout:10,
		}
		config=MemCacheConnect{
			Master:master,
			Slave:slaves,
		}
		file.PutContent(memcacheFile,fmt.Sprintf("%v",config))
		panic("please setting redis config in config/memcached.config file")
	}
	lib.StringToObject(cfgFile, &config)
	var rds Memcach
	return rds.SetPool(NewPool(config))
}

//NewPool 初始化数据库连接
func NewPool(config MemCacheConnect) MemcachePooler{
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool==nil{
		pool=&MemcachePool{}
		pool.Init(config)
	}
	return pool;
}