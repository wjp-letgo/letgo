package binding

import (
	"net/http"
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/web/headerlock"
)

//gzipTextBinding
type gzipTextBinding struct{}

//Render
func(gzipTextBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"application/json; charset=utf-8"})
	headerlock.HeaderMapMutex.RLock()
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(code)
	headerlock.HeaderMapMutex.RUnlock()
	str:=value.(string)
	_,err:=w.Write(lib.GzipData([]byte(str)))
	return err
}