package mysql

import (
	"github.com/wjpxxx/letgo/lib"
)
//MysqlConnect mysql连接配置数据
type MysqlConnect struct {
	Master SlaveDB `json:"master"`
	Slave []SlaveDB `json:"slave"`
}
func (m MysqlConnect)String()string{
	return lib.ObjectToString(m)
}
//Slave 从库配置数据
type SlaveDB struct {
	Name string `json:"name"`
	DatabaseName string `json:"databaseName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
	Charset string `json:"charset"`
	MaxOpenConns int `json:"maxOpenConns"`
	MaxIdleConns int `json:"maxIdleConns"`
	MaxLifetime int `json:"maxLifetime"`
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}
//数据库配置
type DBConfig SlaveDB