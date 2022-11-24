package memcache

import "github.com/wjpxxx/letgo/lib"

//MemCacheConnect 连接配置
type MemCacheConnect struct {
	Master SlaveDB `json:"master"`
	Slave []SlaveDB `json:"slave"`
}
//String 连接配置
func (m MemCacheConnect)String()string{
	return lib.ObjectToString(m)
}

//Slave 从库配置数据
type SlaveDB struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
	MaxIdle int `json:"maxIdle"`
	IdleTimeout int `json:"idleTimeout"`
}