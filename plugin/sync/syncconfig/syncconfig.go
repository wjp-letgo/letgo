package syncconfig

import (
	"github.com/wjpxxx/letgo/lib"
)


//ClientConfig 客户端配置文件
type ClientConfig struct {
	Paths []PathList `json:"paths"`
	Server SyncServer `json:"server"`
}


//String
func (c ClientConfig)String()string{
	return lib.ObjectToString(c)
}
//PathList
type PathList struct{
	LocationPath string `json:"locationPath"`
	RemotePath string `json:"remotePath"`
	Filter []string `json:"filter"`
}
//String
func (c PathList)String()string{
	return lib.ObjectToString(c)
}
//SyncServer
type SyncServer struct{
	Server
	Slave []Server	`json:"slave"`
}

//String
func (s SyncServer)String()string{
	return lib.ObjectToString(s)
}

//Server 服务器信息
type Server struct{
	IP string `json:"ip"`
	Port string `json:"port"`
}
//String
func (s Server)String()string{
	return lib.ObjectToString(s)
}

//FileSyncMessage 同步文件信息
type FileSyncMessage struct {
	LocationPath string `json:"locationPath"`
	RemotePath string `json:"remotePath"`
	RelPath string `json:"relPath"`
	File FileData `json:"file"`
	Slave []Server `json:"slave"`
}
//String
func (f FileSyncMessage)String()string{
	return lib.ObjectToString(f)
}
//FileData
type FileData struct{
	Name string `json:"name"`
	Path string `json:"path"`
	Seek int64 `json:"seek"`
	Size int64 `json:"size"`
	Data []byte `json:"data"`
}
//CmdMessage
type CmdMessage struct {
	Server
	Dir string `json:"dir"`
	Cmd string `json:"cmd"`
	Slave []CmdSlave `json:"slave"`
}

//CmdSlave
type CmdSlave struct{
	Server
	Dir string `json:"dir"`
	Cmd string `json:"cmd"`
}
//CmdResult
type CmdResult struct{
	Server
	Result string `json:"result"`
}
//String
func (f CmdResult)String()string{
	return lib.ObjectToString(f)
}

//MessageResult
type MessageResult struct{
	Success bool `json:"success"`
	Err string `json:"err"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []byte `json:"data"`
	
}

//String
func (f MessageResult)String()string{
	return lib.ObjectToString(f)
}