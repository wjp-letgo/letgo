package api

import (
	"bytes"
	"github.com/wjpxxx/letgo/command"
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/net/rpc"
	"github.com/wjpxxx/letgo/plugin/sync/syncconfig"
)

//Command
type Command struct{
}
//Run
func (c *Command)Run(message syncconfig.CmdMessage, out *syncconfig.MessageResult)error{
	var results map[string]syncconfig.CmdResult=make(map[string]syncconfig.CmdResult)
	rs:=c.doCmd(message)
	if rs!=nil{
		k:=(*rs).Server.IP+":"+(*rs).Server.Port
		results[k]=*rs
	}
	results=c.doSlaveCmd(message,results)
	out.Code=200
	out.Success=true
	out.Msg="成功"
	out.Data=lib.Serialize(results)
	return nil
}

//doCmd
func  (c *Command)doCmd(message syncconfig.CmdMessage) *syncconfig.CmdResult{
	if message.Cmd!=""{
		cmd:=command.New()
		if message.Dir!=""{
			cmd=cmd.Cd(message.Dir)
		}
		cmd.SetCMD(message.Cmd)
		var result bytes.Buffer
		cmd.SetStdout(&result)
		cmd.SetStderr(&result)
		cmd.Run()
		rs:=result.String()
		return &syncconfig.CmdResult{
			Server: syncconfig.Server{
				IP: message.IP,
				Port: message.Port,
			},
			Result: rs,
		}
	}
	return nil
}
//doSlaveCmd
func  (c *Command)doSlaveCmd(message syncconfig.CmdMessage,results map[string]syncconfig.CmdResult)map[string]syncconfig.CmdResult{
	for _,slave:=range message.Slave{
		if slave.Cmd==""{
			continue
		}
		msg:=syncconfig.CmdMessage{
			Server: syncconfig.Server{
				IP: slave.IP,
				Port: slave.Port,
			},
			Dir: slave.Dir,
			Cmd: slave.Cmd,
			Slave:nil,
		}
		client,err:=rpc.NewClient().WithAddress(slave.IP,slave.Port)
		if err!=nil{
			continue
		}
		for{
			var result syncconfig.MessageResult=syncconfig.MessageResult{}
			_,err=client.Call("Command.Run",msg, &result)
			if err!=nil{
				break
			}
			if result.Success {
				var rs map[string]syncconfig.CmdResult
				lib.UnSerialize(result.Data, &rs)
				for k,v:=range rs{
					results[k]=v
				}
				break
			}
			//重发
		}
		client.Close()
	}
	return results
}