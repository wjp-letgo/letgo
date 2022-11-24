package rpc

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//Client
type Client struct{
	address string
	conn net.Conn
	client *rpc.Client
}
//WithAddress 设置地址
func (c *Client)WithAddress(addr ...string)(*Client,error){
	c.address=lib.ResolveAddress(addr)
	var err error
	c.conn,err=net.Dial("tcp", c.address)
	if err!=nil{
		log.DebugPrint("RPC Dial error %v", err)
		return nil,err
	}
	c.client=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c.conn))
	return c,nil
}
//Start 启动
func (c *Client)Start()(*Client,error){
	return c.WithAddress()
}
//Close 关闭连接
func (c *Client)Close(){
	c.client.Close()
	c.conn.Close()
}
//Call 调用
func (c *Client)Call(serviceMethod string, args interface{}, reply interface{})(*Client,error){
	var err error
	err=c.client.Call(serviceMethod,args,reply)
	if err!=nil{
		log.DebugPrint("IP:%s,RPC Call error %v",c.address,err)
		return nil,err
	}
	return c,nil
}
//CallByMessage
func (c *Client)CallByMessage(message RpcMessage)(*Client,error){
	var err error
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c.conn))
	defer client.Close()
	var reply interface{}
	err=client.Call(message.Method,message.Args,&reply)
	if err!=nil{
		log.DebugPrint("RPC CallByMessage error %v",err)
		return nil,err
	}
	if message.Callback!=nil{
		message.Callback(reply)
	}
	return c,nil
}
//NewClient
func NewClient()*Client{
	return &Client{}
}
//RpcMessage
type RpcMessage struct {
	Method string
	Args interface{}
	Callback func(reply interface{})
}