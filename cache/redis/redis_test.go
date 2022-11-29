package redis

import (
	"fmt"
	"testing"

	"github.com/wjp-letgo/letgo/lib"
)


func TestRedis(t *testing.T) {
	rds:=NewRedis().Master()
	i:=lib.InRow{
		"name":"wjp",
		"age":11,
	}
	k:=lib.InRow{}
	//s:=rds.SMembers("kk")
	rds.SAdd("kk",i)
	rds.SPop("kk",&k)
	fmt.Println(k,rds.Set("a",1,3600))
}