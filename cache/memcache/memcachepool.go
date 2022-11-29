package memcache

import (
	"fmt"
	"time"

	"github.com/wjp-letgo/letgo/lib"

	"github.com/bradfitz/gomemcache/memcache"
)

//MemcachePooler 连接池接口
type MemcachePooler interface{
	GetMaster()*memcache.Client
	GetSlave()*memcache.Client
	GetSlaveByName(name string)*memcache.Client
	Init(connect MemCacheConnect)
}

type poolMemcache struct {
	master *memcache.Client
	slave map[string]*memcache.Client
}

//RedisPool 连接池
type MemcachePool struct {
	pool *poolMemcache
}

//open 打开连接池
func (m *MemcachePool) open(connect SlaveDB) *memcache.Client {
	address:=fmt.Sprintf("%s:%s",connect.Host,connect.Port)
	con:=memcache.New(address)
	con.Timeout=time.Duration(connect.IdleTimeout) * time.Second
	con.MaxIdleConns=connect.MaxIdle
	return con
}

//Init 初始化
func (r *MemcachePool)Init(connect MemCacheConnect) {
	master:=r.open(connect.Master)
	if master!=nil{
		slave:=make(map[string]*memcache.Client)
		r.pool=&poolMemcache{
			master:master,
			slave:slave,
		}
	}
	for _,connectSlave:=range connect.Slave{
		slave:=r.open(connectSlave)
		if slave!=nil{
			r.pool.slave[connectSlave.Name]=slave
		}
	}
}

//GetMaster 获得主memcache
func (r *MemcachePool) GetMaster()*memcache.Client{
	return r.pool.master
}
//GetSlave 获得从memcache
func (r *MemcachePool) GetSlave()*memcache.Client{
	slaveCount:=len(r.pool.slave)
		if slaveCount>0{
			slaveIndex:=r.randSlaveName()
			return r.pool.slave[slaveIndex];
		}else{
			return r.pool.master
		}
}
//randSlaveName 随机从库名
func (r *MemcachePool)randSlaveName()string{
	var keys []string
	for k,_:=range r.pool.slave{
		keys=append(keys, k)
	}
	slaveCount:=len(keys)
	slaveIndex:=lib.Rand(0,slaveCount-1,lib.Time())
	return keys[slaveIndex]
}
//GetSlaveByName 获得从memcache通过名称
func (r *MemcachePool)GetSlaveByName(name string)*memcache.Client{
	return r.pool.slave[name]
}