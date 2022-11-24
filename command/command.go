package command

import (
	"io"
	"os"
	"os/exec"
	"runtime"
)

//Command
type Command struct{
	in io.Reader
	out io.Writer
	errOut io.Writer
	cmd *exec.Cmd
	Name string
	Args []string
	pipeCmd *Command
	dir string
}
//newCmd
func (c *Command)newCmd()*exec.Cmd{
	var cmd *exec.Cmd
	if runtime.GOOS=="windows" {
		var cmds []string
		cmds=append([]string{"/C",c.Name}, c.Args...)
		cmd=exec.Command("cmd",cmds...)
	}else{
		var cmds []string
		cmds=append([]string{"-c",c.Name}, c.Args...)
		cmd=exec.Command("/bin/sh",cmds...)
	}
	if c.dir!=""{
		cmd.Dir=c.dir
	}
	//cmd.SysProcAttr=&syscall.SysProcAttr{HideWindow: true,CreationFlags: 0x08000000}
	return cmd
}
//Cd
func (c *Command)Cd(path string)*Command{
	c.dir=path
	return c
}

//Run
func(c *Command)Run(){
	c.cmd=c.newCmd()
	setIn(c)
	if c.pipeCmd==nil{
		if c.out!=nil{
			c.cmd.Stdout=c.out
		}else{
			c.cmd.Stdout=os.Stdout
		}
	}else{
		setOut(c.pipeCmd,c)
	}
	if c.pipeCmd==nil{
		if c.errOut!=nil{
			c.cmd.Stderr=c.errOut
		}else{
			c.cmd.Stderr=os.Stdout
		}
	}else{
		setErr(c.pipeCmd,c)
	}
	start(c)
	wait(c)
}
//setErr
func setErr(c *Command,p *Command){
	if c==nil&&p!=nil{
		if p.out!=nil{
			p.cmd.Stderr=p.out
		}else{
			p.cmd.Stderr=os.Stdout
		}
	}else{
		c.SetStderr(p.errOut)
		setErr(c.pipeCmd,c)
	}
}
//setOut
func setOut(c *Command,p *Command){
	if c==nil&&p!=nil{
		if p.out!=nil{
			p.cmd.Stdout=p.out
		}else{
			p.cmd.Stdout=os.Stdout
		}
	}else{
		c.SetStdout(p.out)
		setOut(c.pipeCmd,c)
	}
}
//setIn
func setIn(c *Command){
	if c.pipeCmd!=nil{
		c.pipeCmd.cmd=c.pipeCmd.newCmd()
		c.pipeCmd.cmd.Stdin=c.StdoutPipe()
		setIn(c.pipeCmd)
	}
}
//start
func start(c *Command){
	c.Start()
	if c.pipeCmd!=nil{
		start(c.pipeCmd)
	}
}

//wait
func wait(c *Command){
	c.Wait()
	if c.pipeCmd!=nil{
		wait(c.pipeCmd)
	}
}

//AddPipe
func (c *Command)AddPipe(cmd *Command){
	c.pipeCmd=cmd
}
//StdoutPipe
func (c *Command)StdoutPipe()io.ReadCloser{
	out,_:=c.cmd.StdoutPipe()
	return out
}
//Start
func (c *Command)Start()error{
	return c.cmd.Start()
}
//Wait
func (c *Command)Wait()error{
	return c.cmd.Wait()
}
//SetStdin
func (c *Command)SetStdin(in io.Reader){
	c.in=in
}
//SetStdout
func (c *Command)SetStdout(out io.Writer){
	c.out=out
}
//SetStderr
func (c *Command)SetStderr(err io.Writer){
	c.errOut=err
}
func (c *Command)SetCMD(name string,args ...string)*Command{
	c.Name=name
	c.Args=args
	return c
}
//New
func New()*Command{
	c:= &Command{
	}
	return c
}