package binding

import (
	"html/template"
	"net/http"
	"github.com/wjpxxx/letgo/web/headerlock"
)

//HTML
type HTML struct {
	Template *template.Template
	Name string
}
//NewHTML
func NewHTML(name string,template *template.Template) Rendering{
	return HTML{
		Name: name,
		Template: template,
	}
}

//Render
func(h HTML)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"text/html; charset=utf-8"})
	headerlock.HeaderMapMutex.RLock()
	w.WriteHeader(code)
	headerlock.HeaderMapMutex.RUnlock()
	if h.Name==""{
		return h.Template.Execute(w,value)
	}
	return h.Template.ExecuteTemplate(w,h.Name,value)
}