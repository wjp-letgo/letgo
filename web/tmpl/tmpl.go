package tmpl

import (
	"html/template"
	"sync"
)

//Tmpl
type Tmpl struct{
	Template *template.Template
	delims Delims
	funcMap template.FuncMap
}
//Delims
type Delims struct{
	Left string
	Right string
}
var onceDo sync.Once
var initTmpl *Tmpl

//GetTmpl
func GetTmpl()*Tmpl{
	onceDo.Do(func(){
		initTmpl=&Tmpl{
			delims: Delims{Left: "{{",Right: "}}"},
			funcMap: template.FuncMap{},
		}
	})
	return initTmpl
}
//SetDelims
func (t *Tmpl)SetDelims(left,right string)*Tmpl{
	t.delims=Delims{Left: left,Right: right}
	return t
}
//SetFuncMap
func (t *Tmpl)SetFuncMap(funcMap template.FuncMap){
	t.funcMap=funcMap
}
//LoadHTMLGlob
func (t *Tmpl)LoadHTMLGlob(pattern string){
	t.Template=template.Must(template.New("").Delims(t.delims.Left,t.delims.Right).Funcs(t.funcMap).ParseGlob(pattern))
}
//LoadHTMLFiles
func (t *Tmpl)LoadHTMLFiles(files ...string){
	t.Template=template.Must(template.New("").Delims(t.delims.Left,t.delims.Right).Funcs(t.funcMap).ParseFiles(files...))
}