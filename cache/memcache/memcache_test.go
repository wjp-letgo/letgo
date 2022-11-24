package memcache

import (
	"fmt"
	"testing")


func TestMemcache(t *testing.T) {
	m:=NewMemcache()
	fmt.Println(m.Set("a",1,3600))
}