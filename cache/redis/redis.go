package redis

import (
	"fmt"
	"sync"

	"github.com/garyburd/redigo/redis"
	"github.com/wjp-letgo/letgo/file"
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/log"
)

// 全局实现者
var pool RedisPooler
var poolLock sync.Mutex

// NewPool 初始化数据库连接
func NewPool(config RedisConnect) RedisPooler {
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool == nil {
		pool = &RedisPool{}
		pool.Init(config)
	}
	return pool
}

// Rediser 操作接口
type Rediser interface {
	Master() Master
	Slave() Master
	SlaveByName(name string) Master
}

// Master 从接口
type Master interface {
	Set(key string, value interface{}, overtime int64) bool
	SetNoFix(key string, value interface{}, overtime int64) bool
	SetNx(key string, value interface{}) bool
	Get(key string, value interface{}) bool
	GetNoFix(key string, value interface{}) bool
	Incr(key string) bool
	Incrby(key string,num int64) bool
	Decr(key string) bool
	Decrby(key string,num int64) bool
	Del(key string) bool
	Ttl(key string) int64
	Expire(key string, overtime int64) bool
	Len(key string) int64
	FlushDB() bool
	Exists(key string) bool
	Keys(key string) []string
	RPush(key string, value ...interface{}) int64
	LPush(key string, value ...interface{}) int64
	LPop(key string, value interface{}) bool
	RPop(key string, value interface{}) bool
	Type(key string) (string, bool)
	Ping() bool
	GetRequirepass() bool
	SetRequirepass(password string) bool
	Select(index int) bool
	HMset(key string, value lib.InRow) bool
	HDel(key string, field ...string) int
	HExists(key string, field string) int
	HGet(key string, field string, value interface{}) bool
	HGetAll(key string) lib.Row
	HLen(key string) int
	HKeys(key string) []string
	HSet(key string, field string, value interface{}) int
	HSetNx(key string, field string, value interface{}) bool
	SAdd(key string, values ...interface{}) int
	SCard(key string) int
	SDiff(keys ...string) [][]byte
	SDiffStore(destination string, keys ...string) int
	SInter(keys ...string) [][]byte
	SInterStore(destination string, keys ...string) int
	SIsMember(key string, value interface{}) bool
	SMembers(key string) [][]byte
	SMove(source, destination string, member interface{}) int
	SPop(key string, value interface{}) bool
	SRandMember(key string, count int) [][]byte
	SRem(key string, members ...interface{}) int
	Push(key string, value interface{}) bool
	Pop(key string, value interface{}) bool
}

// Redis redis对象
type Redis struct {
	pool      RedisPooler
	isMaster  bool
	slaveName string
}

// Master 主redis
func (r *Redis) Master() Master {
	r.isMaster = true
	return r
}

// Slave 从redis
func (r *Redis) Slave() Master {
	r.isMaster = false
	r.slaveName = ""
	return r
}

// SlaveByName 从redis 通过名称
func (r *Redis) SlaveByName(name string) Master {
	r.isMaster = false
	r.slaveName = name
	return r
}

// getRedis 获得redis
func (r *Redis) getRedis() *redis.Pool {
	if r.isMaster {
		return r.pool.GetMaster()
	} else {
		if r.slaveName == "" {
			return r.pool.GetSlave()
		} else {
			return r.pool.GetSlaveByName(r.slaveName)
		}

	}
}

// SetPool 设置池
func (r *Redis) SetPool(pooler RedisPooler) Rediser {
	r.pool = pooler
	return r
}

// Set set操作
func (r *Redis) Set(key string, value interface{}, overtime int64) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	if overtime > -1 {
		_, err := rds.Do("SET", key, lib.Serialize(value), "EX", overtime)
		if err != nil {
			log.DebugPrint("redis set fail: %s", err.Error())
			return false
		}
	} else {
		_, err := rds.Do("SET", key, lib.Serialize(value))
		if err != nil {
			log.DebugPrint("redis set fail: %s", err.Error())
			return false
		}
	}
	return true
}

// Set set操作
func (r *Redis) SetNoFix(key string, value interface{}, overtime int64) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	if overtime > -1 {
		_, err := rds.Do("SET", key, lib.SerializeNoFix(value), "EX", overtime)
		if err != nil {
			log.DebugPrint("redis set fail: %s", err.Error())
			return false
		}
	} else {
		_, err := rds.Do("SET", key, lib.SerializeNoFix(value))
		if err != nil {
			log.DebugPrint("redis set fail: %s", err.Error())
			return false
		}
	}
	return true
}

// Incr  操作
func (r *Redis) Incr(key string) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	_, err := rds.Do("INCR", key)
	if err != nil {
		log.DebugPrint("redis Incr fail: %s", err.Error())
		return false
	}
	return true
}

// Incrby  操作
func (r *Redis) Incrby(key string,num int64) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	_, err := rds.Do("INCRBY", key,num)
	if err != nil {
		log.DebugPrint("redis Incrby fail: %s", err.Error())
		return false
	}
	return true
}

// Decr 操作
func (r *Redis) Decr(key string) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	_, err := rds.Do("DECR", key)
	if err != nil {
		log.DebugPrint("redis Decr fail: %s", err.Error())
		return false
	}
	return true
}

// Decrby  操作
func (r *Redis) Decrby(key string,num int64) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	_, err := rds.Do("DECRBY", key,num)
	if err != nil {
		log.DebugPrint("redis Decrby fail: %s", err.Error())
		return false
	}
	return true
}

// SetNx SetNx 操作
func (r *Redis) SetNx(key string, value interface{}) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int64(rds.Do("SETNX", key, value))
	if err != nil {
		log.DebugPrint("redis setNx fail: %s", err.Error())
		return false
	}
	if v == 0 {
		return false
	}
	return true
}

// Get Get操作
func (r *Redis) Get(key string, value interface{}) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	tv, err := rds.Do("GET", key)
	if err != nil {
		log.DebugPrint("redis get fail: %s", err.Error())
		return false
	}
	if tv != nil {
		v := lib.Data{Value: tv}
		lib.UnSerialize(v.ArrayByte(), value)
	}
	return true
}

// Get Get操作
func (r *Redis) GetNoFix(key string, value interface{}) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	tv, err := rds.Do("GET", key)
	if err != nil {
		log.DebugPrint("redis get fail: %s", err.Error())
		return false
	}
	log.DebugPrint("key:%s,getnofix:%v", key, tv)
	if tv != nil {
		v := lib.Data{Value: tv}
		lib.UnSerializeNoFix(v.ArrayByte(), value)
	}
	return true
}

// Del Del操作
func (r *Redis) Del(key string) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	_, err := rds.Do("DEL", key)
	if err != nil {
		log.DebugPrint("redis del fail: %s", err.Error())
		return false
	}
	return true
}

// Ttl Ttl操作
func (r *Redis) Ttl(key string) int64 {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int64(rds.Do("TTL", key))
	if err != nil {
		log.DebugPrint("redis ttl fail: %s", err.Error())
		return -1
	}
	return v
}

// Expire Expire操作
func (r *Redis) Expire(key string, overtime int64) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	_, err := rds.Do("EXPIRE", key, overtime)
	if err != nil {
		log.DebugPrint("redis expire fail: %s", err.Error())
		return false
	}
	return true
}

// Len Len操作
func (r *Redis) Len(key string) int64 {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int64(rds.Do("LLEN", key))
	if err != nil {
		log.DebugPrint("redis llen fail: %s", err.Error())
		return -1
	}
	return v
}

// FlushDB FlushDB操作
func (r *Redis) FlushDB() bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	_, err := rds.Do("FLUSHDB")
	if err != nil {
		log.DebugPrint("redis flushdb fail: %s", err.Error())
		return false
	}
	return true
}

// Exists Exists操作
func (r *Redis) Exists(key string) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("EXISTS", key))
	if err != nil {
		log.DebugPrint("redis exists fail: %s", err.Error())
		return false
	}
	if v == 1 {
		return true
	} else {
		return false
	}
}

// Keys Keys操作
func (r *Redis) Keys(key string) []string {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Strings(rds.Do("KEYS", key))
	if err != nil {
		log.DebugPrint("redis keys fail: %s", err.Error())
		return nil
	}
	return v
}

// RPush RPush操作
func (r *Redis) RPush(key string, value ...interface{}) int64 {
	var arg []interface{}
	arg = append(arg, key)
	for _, d := range value {
		arg = append(arg, lib.Serialize(d))
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int64(rds.Do("Rpush", arg...))
	if err != nil {
		log.DebugPrint("redis rpush fail: %s", err.Error())
		return -1
	}
	return v
}

// LPush LPush操作
func (r *Redis) LPush(key string, value ...interface{}) int64 {
	var arg []interface{}
	arg = append(arg, key)
	for _, d := range value {
		arg = append(arg, lib.Serialize(d))
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int64(rds.Do("Lpush", arg...))
	if err != nil {
		log.DebugPrint("redis lpush fail: %s", err.Error())
		return -1
	}
	return v
}

// LPop LPop操作
func (r *Redis) LPop(key string, value interface{}) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	tv, err := rds.Do("Lpop", key)
	if err != nil {
		log.DebugPrint("redis lpop fail: %s", err.Error())
		return false
	}
	if tv != nil {
		v := lib.Data{Value: tv}
		lib.UnSerialize(v.ArrayByte(), value)
		return true
	}
	return false
}

// RPop RPop操作
func (r *Redis) RPop(key string, value interface{}) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	tv, err := rds.Do("Rpop", key)
	if err != nil {
		log.DebugPrint("redis rpop fail: %s", err.Error())
		return false
	}
	if tv != nil {
		v := lib.Data{Value: tv}
		lib.UnSerialize(v.ArrayByte(), value)
		return true
	}
	return false
}

// Push Push操作
func (r *Redis) Push(key string, value interface{}) bool {
	i := r.LPush(key, value)
	if i > -1 {
		return true
	}
	return false
}

// Push Push操作
func (r *Redis) Pop(key string, value interface{}) bool {
	return r.RPop(key, value)
}

// Type Type操作
func (r *Redis) Type(key string) (string, bool) {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.String(rds.Do("TYPE", key))
	if err != nil {
		log.DebugPrint("redis type fail: %s", err.Error())
		return "", false
	}
	return v, true
}

// Ping Ping操作
func (r *Redis) Ping() bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.String(rds.Do("PING"))
	if err != nil {
		log.DebugPrint("redis ping fail: %s", err.Error())
		return false
	}
	if v == "PONG" {
		return true
	} else {
		return false
	}
}

// GetRequirepass GetRequirepass操作
func (r *Redis) GetRequirepass() bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Strings(rds.Do("CONFIG", "get", "requirepass"))
	if err != nil {
		log.DebugPrint("redis config get requirepass fail: %s", err.Error())
		return false
	}
	//fmt.Println(v)
	if v[1] == "" {
		return false
	} else {
		return true
	}
}

// SetRequirepass SetRequirepass操作
func (r *Redis) SetRequirepass(password string) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.String(rds.Do("CONFIG", "set", "requirepass", password))
	if err != nil {
		log.DebugPrint("redis config set requirepass fail: %s", err.Error())
		return false
	}
	if v == "OK" {
		return true
	} else {
		return false
	}
}

// Select 选择数据库
func (r *Redis) Select(index int) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.String(rds.Do("SELECT", index))
	if err != nil {
		log.DebugPrint("redis select fail: %s", err.Error())
		return false
	}
	if v == "OK" {
		return true
	} else {
		return false
	}
}

// HMset HMset操作
func (r *Redis) HMset(key string, value lib.InRow) bool {
	var arg []interface{}
	arg = append(arg, key)
	for k, iv := range value {
		arg = append(arg, k)
		arg = append(arg, lib.Serialize(iv))
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.String(rds.Do("HMSET", arg...))
	if err != nil {
		log.DebugPrint("redis hmset fail: %s", err.Error())
		return false
	}
	if v == "OK" {
		return true
	} else {
		return false
	}
}

// HDel HDel操作
func (r *Redis) HDel(key string, field ...string) int {
	var args []interface{}
	args = append(args, key)
	for _, v := range field {
		args = append(args, v)
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("HDEL", args...))
	if err != nil {
		log.DebugPrint("redis hdel fail: %s", err.Error())
		return 0
	}
	return v
}

// HExists HExists操作
func (r *Redis) HExists(key string, field string) int {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("HEXISTS", key, field))
	if err != nil {
		log.DebugPrint("redis hexists fail: %s", err.Error())
		return -1
	}
	return v
}

// HGet HGet操作
func (r *Redis) HGet(key string, field string, value interface{}) bool {
	crds := r.getRedis()
	if crds == nil {
		log.DebugPrint("redis hget fail: pool is nil")
		return false
	}
	rds := crds.Get()
	defer rds.Close()
	tv, err := rds.Do("HGET", key, field)
	if err != nil {
		log.DebugPrint("redis hget fail: %s", err.Error())
		return false
	}
	if tv != nil {
		v := lib.Data{Value: tv}
		lib.UnSerialize(v.ArrayByte(), value)
		return true
	}
	return false
}

// HGetAll HGetAll操作
func (r *Redis) HGetAll(key string) lib.Row {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.ByteSlices(rds.Do("HGETALL", key))
	if err != nil {
		log.DebugPrint("redis hgetall fail: %s", err.Error())
		return nil
	}
	value := make(lib.Row)
	for i := 0; i < len(v); i = i + 2 {
		vl := &lib.Data{}
		value[string(v[i])] = vl.Set(v[i+1])
	}
	return value
}

// HLen HLen操作
func (r *Redis) HLen(key string) int {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("HLEN", key))
	if err != nil {
		log.DebugPrint("redis hlen fail: %s", err.Error())
		return -1
	}
	return v
}

// HKeys HKeys操作
func (r *Redis) HKeys(key string) []string {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Strings(rds.Do("HKEYS", key))
	if err != nil {
		log.DebugPrint("redis hkeys fail: %s", err.Error())
		return nil
	}
	return v
}

// HSet HSet操作
func (r *Redis) HSet(key string, field string, value interface{}) int {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("HSET", key, field, lib.Serialize(value)))
	if err != nil {
		log.DebugPrint("redis hset fail: %s", err.Error())
		return -1
	}
	return v
}

// HSetNx HSetNx操作
func (r *Redis) HSetNx(key string, field string, value interface{}) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("HSETNX", key, field, lib.Serialize(value)))
	if err != nil {
		log.DebugPrint("redis hsetnx fail: %s", err.Error())
		return false
	}
	if v == 1 {
		return true
	}
	return false
}

// SAdd SAdd 操作
func (r *Redis) SAdd(key string, values ...interface{}) int {
	var args []interface{}
	args = append(args, key)
	for _, v := range values {
		args = append(args, lib.Serialize(v))
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("SADD", args...))
	if err != nil {
		log.DebugPrint("redis sadd fail: %s", err.Error())
		return -1
	}
	return v
}

// SCard SCard 操作
func (r *Redis) SCard(key string) int {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("SCARD", key))
	if err != nil {
		log.DebugPrint("redis scard fail: %s", err.Error())
		return -1
	}
	return v
}

// SDiff SDiff 操作命令返回第一个集合与其他集合之间的差异
func (r *Redis) SDiff(keys ...string) [][]byte {
	var args []interface{}
	for _, k := range keys {
		args = append(args, k)
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.ByteSlices(rds.Do("SDIFF", args...))
	if err != nil {
		log.DebugPrint("redis sdiff fail: %s", err.Error())
		return nil
	}
	return v
}

// SDiffStore SDiffStore 命令将给定集合之间的差集存储在指定的集合中
func (r *Redis) SDiffStore(destination string, keys ...string) int {
	var args []interface{}
	args = append(args, destination)
	for _, k := range keys {
		args = append(args, k)
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("SDIFFSTORE", args...))
	if err != nil {
		log.DebugPrint("redis sdiffstore fail: %s", err.Error())
		return -1
	}
	return v
}

// SInter SInter 操作
func (r *Redis) SInter(keys ...string) [][]byte {
	var args []interface{}
	for _, k := range keys {
		args = append(args, k)
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.ByteSlices(rds.Do("SINTER", args...))
	if err != nil {
		log.DebugPrint("redis sinter fail: %s", err.Error())
		return nil
	}
	return v
}

// SInterStore SInterStore 操作
func (r *Redis) SInterStore(destination string, keys ...string) int {
	var args []interface{}
	args = append(args, destination)
	for _, k := range keys {
		args = append(args, k)
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("SINTERSTORE", args...))
	if err != nil {
		log.DebugPrint("redis sinterstore fail: %s", err.Error())
		return -1
	}
	return v
}

// SIsMember SIsMember 命令判断成员元素是否是集合的成员。
func (r *Redis) SIsMember(key string, value interface{}) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("SISMEMBER", key, lib.Serialize(value)))
	if err != nil {
		log.DebugPrint("redis sismember fail: %s", err.Error())
		return false
	}
	if v == 1 {
		return true
	}
	return false
}

// SMembers SMembers 命令返回集合中的所有的成员。 不存在的集合 key 被视为空集合。
func (r *Redis) SMembers(key string) [][]byte {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.ByteSlices(rds.Do("SMEMBERS", key))
	if err != nil {
		log.DebugPrint("redis smembers fail: %s", err.Error())
		return nil
	}
	return v
}

// SMove SMove 将 member 元素从 source 集合移动到 destination 集合
func (r *Redis) SMove(source, destination string, member interface{}) int {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("SMOVE", source, destination, lib.Serialize(member)))
	if err != nil {
		log.DebugPrint("redis smove fail: %s", err.Error())
		return -1
	}
	return v
}

// SPop SPop 移除并返回集合中的一个随机元素
func (r *Redis) SPop(key string, value interface{}) bool {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Bytes(rds.Do("SPOP", key))
	if err != nil {
		log.DebugPrint("redis spop fail: %s", err.Error())
		return false
	}
	lib.UnSerialize(v, value)
	return true
}

// SRandMember SRandMember 返回集合中一个或多个随机数
func (r *Redis) SRandMember(key string, count int) [][]byte {
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.ByteSlices(rds.Do("SRANDMEMBER", key, count))
	if err != nil {
		log.DebugPrint("redis srandmember fail: %s", err.Error())
		return nil
	}
	return v
}

// SRem SRem 移除集合中一个或多个成员
func (r *Redis) SRem(key string, members ...interface{}) int {
	var args []interface{}
	args = append(args, key)
	for _, k := range members {
		args = append(args, lib.Serialize(k))
	}
	rds := r.getRedis().Get()
	defer rds.Close()
	v, err := redis.Int(rds.Do("SREM", args...))
	if err != nil {
		log.DebugPrint("redis srem fail: %s", err.Error())
		return -1
	}
	return v
}

// NewRedis 新建一个redis
func NewRedis() Rediser {
	//log.DebugPrint("#############################创建1")
	redisFile := "config/redis.config"
	cfgFile := file.GetContent(redisFile)
	var config RedisConnect
	if cfgFile == "" {
		var slaves []SlaveDB = make([]SlaveDB, 1)
		master := SlaveDB{
			Name:        "name",
			Db:          0,
			Password:    "password",
			Host:        "127.0.0.1",
			Port:        "6379",
			MaxIdle:     20,
			IdleTimeout: 10,
			MaxActive:   100,
		}
		config = RedisConnect{
			Master: master,
			Slave:  slaves,
		}
		file.PutContent(redisFile, fmt.Sprintf("%v", config))
		log.PanicPrint("please setting redis config in config/redis.config file")
	}
	lib.StringToObject(cfgFile, &config)
	var rds Redis
	return rds.SetPool(NewPool(config))
}

// NewRedisByConnect
func NewRedisByConnect(config RedisConnect) Rediser {
	//log.DebugPrint("#############################创建2")
	var rds Redis
	return rds.SetPool(NewPool(config))
}
