package binding

import (
	"net/http"

	"github.com/wjp-letgo/letgo/web/headerlock"
)

//Binding
type Binding interface {
	Name() string
	Bind(*http.Request,[]byte,interface{})error
}
//Render
type Rendering interface{
	Render(int,http.ResponseWriter,interface{})error
}
const (
	MIMEJSON              = "application/json"
	MIMEJSONUTF8          = "application/json;charset=utf-8"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEYAML              = "application/x-yaml"
	MIMETEXT              = "text/html"
)
var (
	JSON =jsonBinding{}
	GZIPJSON=gzipJsonBinding{}
	XML=xmlBinding{}
	YAML=yamlBinding{}
	JSONP=jsonpBinding{}
	GZIPJSONP=gzipJsonpBinding{}
	TEXT=textBinding{}
	GZIPTEXT=gzipTextBinding{}
)
func NewBind(contentType string)Binding{
	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEJSONUTF8:
		return JSON
	case MIMEXML,MIMEXML2:
		return XML
	case MIMEYAML:
		return YAML
	default:
		panic("unknown type")
	}
}

func NewRender(contentType string)Rendering{
	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEJSONUTF8:
		return JSON
	case MIMEXML,MIMEXML2:
		return XML
	case MIMEYAML:
		return YAML
	case MIMETEXT:
		return TEXT
	default:
		panic("unknown type")
	}
}
//writeContentType
func writeContentType(w http.ResponseWriter, value []string) {
	headerlock.HeaderMapMutex.RLock()
	header:=w.Header()
	if v:=header["Content-Type"];len(v)==0{
		header["Content-Type"] = value
	}
	headerlock.HeaderMapMutex.RUnlock()
}