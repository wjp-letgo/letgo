package mssql

import (
	"github.com/wjpxxx/letgo/lib"
)

//MsSqlConnect mssql连接配置数据
type MsSqlConnect struct {
	Master DBConfig `json:"master"`
	Slave []DBConfig `json:"slave"`
}
func (m MsSqlConnect)String()string{
	return lib.ObjectToString(m)
}
//DBConfig 从库配置数据
type DBConfig struct {
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