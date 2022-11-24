package filecache

import (
	"fmt"
	"testing"
)


func TestFileCache(t *testing.T) {
	f:=NewFileCache()
	f.Set("b",1, 3600)
	var i int
	f.Get("a", &i)
	//f.FlushDB()
	f.Del("b")
	fmt.Println(i)
}
