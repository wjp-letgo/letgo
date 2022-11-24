package cache

import (
	"fmt"
	"testing")


func TestCache(t *testing.T){
	r:=NewCache("redis")
	r.Set("a",1,3600)
	var i int 
	r.Get("a", &i)
	f:=NewCache("file")
	f.Set("a",1,3000)
	var j int
	f.Get("a", &j)
	fmt.Println(i)
}