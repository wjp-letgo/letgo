package syncclient

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/net/rpc"
	"github.com/wjpxxx/letgo/plugin/sync/syncconfig"
)

//CommandSync
type CommandSync struct{
	config syncconfig.ClientConfig
}

//Run
func (c *CommandSync)Run(values ...interface{})interface{}{
	client,err:=rpc.NewClient().WithAddress(c.config.Server.IP,c.config.Server.Port)
	if err!=nil{
		return nil
	}
	defer client.Close()
	return c.do(client,values...)
}

//NewCommandSync
func NewCommandSync()*CommandSync{
	return &CommandSync{
		config:config,
	}
}
//NewCommandSyncByConfig
func NewCommandSyncByConfig(server syncconfig.SyncServer)*CommandSync{
	c:=syncconfig.ClientConfig{}
	c.Paths=config.Paths
	c.Server=server
	return &CommandSync{config:c}
}
//do
func (c *CommandSync) do(client *rpc.Client,values ...interface{})map[string]syncconfig.CmdResult{
	var dir string
	var cmd string
	var ip string
	if len(values)==1{
		cmd=values[0].(string)
	}else if len(values)==2{
		dir=values[0].(string)
		cmd=values[1].(string)
	}else if len(values)==3{
		dir=values[0].(string)
		cmd=values[1].(string)
		ip=values[2].(string)
	}
	message:=c.packedCmd(dir,cmd,ip)
	return c.rpcCall(client,message)
}

//rpcCall
func (c *CommandSync)rpcCall(client *rpc.Client,message syncconfig.CmdMessage)map[string]syncconfig.CmdResult{
	for {
		var result syncconfig.MessageResult=syncconfig.MessageResult{}
		client.Call("Command.Run", message,&result)
		if result.Success {
			var rs map[string]syncconfig.CmdResult
			lib.UnSerialize(result.Data, &rs)
			return rs
		}
	}
}

//packed 打包
func (c *CommandSync) packedCmd(dir,cmd,ip string)syncconfig.CmdMessage{
	if ip!=""{
		if ip==c.config.Server.IP{
			//如果有传IP,则只有IP的服务器执行
			return syncconfig.CmdMessage{
				Server: syncconfig.Server{
					IP: c.config.Server.IP,
					Port: c.config.Server.Port,
				},
				Dir: dir,
				Cmd: cmd,
				Slave:c.cmdSlave(dir,cmd,ip),
			}
		}else{
			//其他台不执行
			return syncconfig.CmdMessage{
				Server: syncconfig.Server{
					IP: c.config.Server.IP,
					Port: c.config.Server.Port,
				},
				Dir: "",
				Cmd: "",
				Slave:c.cmdSlave(dir,cmd,ip),
			}
		}
		
	}else{
		return syncconfig.CmdMessage{
			Server: syncconfig.Server{
				IP: c.config.Server.IP,
				Port: c.config.Server.Port,
			},
			Dir: dir,
			Cmd: cmd,
			Slave:c.cmdSlave(dir,cmd,ip),
		}
	}
	
}
//cmdSlave
func (c *CommandSync) cmdSlave(dir,cmd,ip string)[]syncconfig.CmdSlave{
	var slaves []syncconfig.CmdSlave
	for _,slave:=range c.config.Server.Slave{
		if ip!=""{
			if ip==slave.IP{
				//改台IP执行
				s:=syncconfig.CmdSlave{
					Server: syncconfig.Server{
						IP: slave.IP,
						Port: slave.Port,
					},
					Dir: dir,
					Cmd: cmd,
				}
				slaves=append(slaves, s)
			}else{
				s:=syncconfig.CmdSlave{
					Server: syncconfig.Server{
						IP: slave.IP,
						Port: slave.Port,
					},
					Dir: "",
					Cmd: "",
				}
				slaves=append(slaves, s)
			}
		}else{
			s:=syncconfig.CmdSlave{
				Server: syncconfig.Server{
					IP: slave.IP,
					Port: slave.Port,
				},
				Dir: dir,
				Cmd: cmd,
			}
			slaves=append(slaves, s)
		}
	}
	return slaves
}

