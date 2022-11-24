package binding

import (
	"encoding/json"
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/web/headerlock"
	"errors"
	"net/http"
	"bytes"
)

//gzipJsonBinding
type gzipJsonBinding struct{}
//Name
func(gzipJsonBinding)Name()string{
	return "gzipjson"
}
//Bind
func(gzipJsonBinding)Bind(req *http.Request,body []byte,value interface{}) error{
	if req==nil||body==nil{
		return errors.New("error request")
	}
	decoder:=json.NewDecoder(bytes.NewReader(body))
	if err:=decoder.Decode(value);err!=nil{
		return err
	}
	return nil
}
//Render
func(gzipJsonBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"application/json; charset=utf-8"})
	headerlock.HeaderMapMutex.RLock()
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(code)
	headerlock.HeaderMapMutex.RUnlock()
	jsonData,err:=json.Marshal(value)
	if err!=nil{
		return err
	}
	_,err=w.Write(lib.GzipData(jsonData))
	return err
}

//gzipJsonpBinding
type gzipJsonpBinding struct{}

//Render
func(gzipJsonpBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"application/javascript; charset=utf-8"})
	headerlock.HeaderMapMutex.RLock()
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(code)
	headerlock.HeaderMapMutex.RUnlock()
	jsonData,err:=json.Marshal(value)
	if err!=nil{
		return err
	}
	_,err=w.Write(lib.GzipData(jsonData))
	return err
}