package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/log"
)

//RedisPooler 连接池接口
type RedisPooler interface{
	GetMaster()*redis.Pool
	GetSlave()*redis.Pool
	GetSlaveByName(name string)*redis.Pool
	Init(connect RedisConnect)
}
type poolRedis struct {
	master *redis.Pool
	slave map[string]*redis.Pool
}
//RedisPool 连接池
type RedisPool struct {
	pool *poolRedis
}
//open 打开连接池
func (r *RedisPool) open(connect SlaveDB) *redis.Pool {
	return &redis.Pool{
		MaxIdle: connect.MaxIdle,
		Wait: true,
		MaxActive: connect.MaxActive,
		IdleTimeout: time.Duration(connect.IdleTimeout) * time.Second,
		Dial: func () (redis.Conn, error) {
			address:=fmt.Sprintf("%s:%s",connect.Host,connect.Port)
			con,err:=redis.Dial("tcp", address)
			//log.DebugPrint("=========================建立连接:%d",r.pool.master.ActiveCount())
			if err!=nil{
				log.PanicPrint("open connect fail %s",err.Error())
				return nil,err
			}
			if connect.Password!=""{
				if _, err := con.Do("AUTH", connect.Password); err != nil {
					con.Close()
					if err!=nil{
						log.PanicPrint("AUTH fail %s",err.Error())
					}
					return nil, err
				}
			}
			if _, err := con.Do("SELECT", connect.Db); err != nil {
				con.Close()
				if err!=nil{
					log.PanicPrint("SELECT fail %s",err.Error())
				}
				return nil, err
			}
			return con, nil
		},
	}
}
//Init 初始化
func (r *RedisPool)Init(connect RedisConnect) {
	master:=r.open(connect.Master)
	if master!=nil{
		slave:=make(map[string]*redis.Pool)
		r.pool=&poolRedis{
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
//GetMaster 获得主redis
func (r *RedisPool) GetMaster()*redis.Pool{
	return r.pool.master
}
//GetSlave 获得从redis
func (r *RedisPool) GetSlave()*redis.Pool{
	slaveCount:=len(r.pool.slave)
		if slaveCount>0{
			slaveIndex:=r.randSlaveName()
			return r.pool.slave[slaveIndex];
		}else{
			return r.pool.master
		}
}
//randSlaveName 随机从库名
func (r *RedisPool)randSlaveName()string{
	var keys []string
	for k,_:=range r.pool.slave{
		keys=append(keys, k)
	}
	slaveCount:=len(keys)
	slaveIndex:=lib.Rand(0,slaveCount-1,lib.Time())
	return keys[slaveIndex]
}
//GetSlaveByName 获得从redis通过名称
func (r *RedisPool)GetSlaveByName(name string)*redis.Pool{
	return r.pool.slave[name]
}