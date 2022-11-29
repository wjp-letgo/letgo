package binding

import (
	"net/http"

	"github.com/wjp-letgo/letgo/web/headerlock"
)

//textBinding
type textBinding struct{}

//Render
func(textBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"text/html; charset=utf-8"})
	headerlock.HeaderMapMutex.RLock()
	w.WriteHeader(code)
	headerlock.HeaderMapMutex.RUnlock()
	str:=value.(string)
	_,err:=w.Write([]byte(str))
	return err
}