package httpclient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/wjp-letgo/letgo/lib"
)

func TestHttp(t *testing.T){
	c:=&HttpClient{}
	fmt.Println(c.WithRequestBefore(func(req *http.Request){
		fmt.Println("dddd")
	}).Post("http://api-www.yutang.cn/api/Login/getSiteInfo",lib.InRow{
		"@a":"httpclient.go",
		"c":2,
	}).Body())
	SaveRemoteFile("https://cdn2.jianshu.io/assets/default_avatar/10-e691107df16746d4a9f3fe9496fd1848.jpg", "a.jpg")
}
