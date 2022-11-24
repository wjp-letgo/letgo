package binding

import (
	"encoding/json"
	"errors"
	"net/http"
	"bytes"
	"github.com/wjpxxx/letgo/web/headerlock"
)

//jsonBinding
type jsonBinding struct{}
//Name
func(jsonBinding)Name()string{
	return "json"
}
//Bind
func(jsonBinding)Bind(req *http.Request,body []byte,value interface{}) error{
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
func(jsonBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"application/json; charset=utf-8"})
	headerlock.HeaderMapMutex.RLock()
	w.WriteHeader(code)
	headerlock.HeaderMapMutex.RUnlock()
	jsonData,err:=json.Marshal(value)
	if err!=nil{
		return err
	}
	_,err=w.Write(jsonData)
	return err
}

//jsonpBinding
type jsonpBinding struct{}

//Render
func(jsonpBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"application/javascript; charset=utf-8"})
	headerlock.HeaderMapMutex.RLock()
	w.WriteHeader(code)
	headerlock.HeaderMapMutex.RUnlock()
	jsonData,err:=json.Marshal(value)
	if err!=nil{
		return err
	}
	_,err=w.Write(jsonData)
	return err
}