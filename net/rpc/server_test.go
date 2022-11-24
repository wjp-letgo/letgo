package rpc

import (
	"fmt"
	"testing"
	"time"
)

//Hello
type Hello struct{

}

func(h *Hello)Say(in string,out *string) error{
	*out=in+":123123"
	return nil
}

func TestServer(t *testing.T){
	s:=NewServer()
	//s.RegisterName("Hello",new(Hello))
	s.Register(new(Hello))
	go func(){
		for{
			time.Sleep(10*time.Second)
			var reply string
			//NewClient().WithAddress("127.0.0.1","8080").Call("Hello.Say","nihao",&reply).Close()
			NewClient().Start().Call("Hello.Say","nihao",&reply).Close()
			fmt.Println(reply)
			rm:=RpcMessage{
				Method: "Hello.Say",
				Args: "rpc message",
				Callback: func(a interface{}){
					fmt.Println(a.(string))
				},
			}
			NewClient().Start().CallByMessage(rm).Close()
		}
	}()
	s.Run()
}