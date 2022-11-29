package redis

import "github.com/wjp-letgo/letgo/lib"

//RedisConnect 连接配置
type RedisConnect struct {
	Master SlaveDB   `json:"master"`
	Slave  []SlaveDB `json:"slave"`
}

//String 连接配置
func (m RedisConnect) String() string {
	return lib.ObjectToString(m)
}

//Slave 从库配置数据
type SlaveDB struct {
	Name        string `json:"name"`
	Db          int    `json:"db"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	MaxIdle     int    `json:"maxIdle"`
	IdleTimeout int    `json:"idleTimeout"`
	MaxActive   int    `json:"maxActive"`
}
