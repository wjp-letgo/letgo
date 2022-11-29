package rpc

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"

	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/log"
)

//Server
type Server struct {

}

//Run 启动服务
func (s *Server)Run(addr ...string)error {
	address:=lib.ResolveAddress(addr)
	listenner,err:=net.Listen("tcp", address)
	if err!=nil{
		log.DebugPrint("RPC listen error %v", err)
		return err
	}
	for{
		conn,err:=listenner.Accept()
		if err!=nil{
			log.DebugPrint("RPC Accept error %v",err)
			return err
		}
		log.DebugPrint("客户端:%s,连接上来了",conn.RemoteAddr().String())
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
//RegisterName 注册
func (s *Server)RegisterName(name string,server interface{}) *Server{
	err:= rpc.RegisterName(name,server)
	if err!=nil{
		log.DebugPrint("rpc register name error %v",err)
		return nil
	}
	return s
}

//Register 注册
func (s *Server)Register(server interface{}) *Server{
	err:= rpc.Register(server)
	if err!=nil{
		log.DebugPrint("rpc register error %v",err)
		return nil
	}
	return s
}
var initServer *Server
var once sync.Once
//NewClient
func NewServer()*Server{
	once.Do(func(){
		initServer=&Server{}
	})
	return initServer
}