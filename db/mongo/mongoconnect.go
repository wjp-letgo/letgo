package mongo

import "github.com/wjp-letgo/letgo/lib"

//MongoConnect
type MongoConnect struct {
	Name           string  `json:"name"`
	UserName       string  `json:"userName"`
	Password       string  `json:"password"`
	Hosts          []Host  `json:"hosts"`
	Database       string  `json:"database"`
	Option         Options `json:"option"`
	ConnectTimeout int     `json:"connectTimeout"`
	ExecuteTimeout int     `json:"executeTimeout"`
}

//String
func (m MongoConnect) String() string {
	return lib.ObjectToString(m)
}

//Host
type Host struct {
	Hst  string `json:"host"`
	Port string `json:"port"`
}

//Options
type Options struct {
	ReplicaSet       string `json:"replicaSet"`
	SlaveOk          bool   `json:"slaveOk"`
	Safe             bool   `json:"safe"`
	WtimeoutMS       int64  `json:"wtimeoutMS"`
	ConnectTimeoutMS int64  `json:"connectTimeoutMS"`
	SocketTimeoutMS  int64  `json:"socketTimeoutMS"`
	MaxPoolSize      int    `json:"maxPoolSize"`
	MinPoolSize      int    `json:"minPoolSize"`
	MaxIdleTimeMS    int64  `json:"maxIdleTimeMS"`
}
