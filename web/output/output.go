package output

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/web/binding"
	"github.com/wjp-letgo/letgo/web/headerlock"
	"github.com/wjp-letgo/letgo/web/input"

	//"github.com/wjp-letgo/letgo/log"
	"html/template"
	"net/http"
)

//Output
type Output struct {
	writer http.ResponseWriter
	in     *input.Input
	templ  *template.Template
	status int
}

//Init 初始化
func (o *Output) Init(writer http.ResponseWriter, in *input.Input, templ *template.Template) {
	o.writer = writer
	o.in = in
	o.templ = templ
	o.status = 0
}

//Header 设置头
func (o *Output) Header(key, value string) {
	headerlock.HeaderMapMutex.Lock()
	if o != nil && o.writer != nil && o.writer.Header() != nil {
		o.writer.Header().Set(key, value)
	}
	headerlock.HeaderMapMutex.Unlock()
}

//JSON
func (o *Output) JSON(code int, value interface{}) error {
	return o.Render(code, value, binding.JSON)
}

//GZIPJSON
func (o *Output) GZIPJSON(code int, value interface{}) error {
	return o.Render(code, value, binding.GZIPJSON)
}

//JSONOK
func (o *Output) JSONOK(code int, message string) error {
	return o.JSON(code, lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      message,
		"err":      "",
		"sub_code": "success",
	})
}
//GZIPJSONOK
func (o *Output) GZIPJSONOK(code int, message string) error {
	return o.GZIPJSON(code, lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      message,
		"err":      "",
		"sub_code": "success",
	})
}
//JSONERROR
func (o *Output) JSONERROR(code int, message, subCode string) error {
	return o.JSON(code, lib.InRow{
		"code":     0,
		"success":  false,
		"msg":      "",
		"err":      message,
		"sub_code": subCode,
	})
}
//GZIPJSONERROR
func (o *Output) GZIPJSONERROR(code int, message, subCode string) error {
	return o.GZIPJSON(code, lib.InRow{
		"code":     0,
		"success":  false,
		"msg":      "",
		"err":      message,
		"sub_code": subCode,
	})
}
//JSONFail
func (o *Output) JSONFail(code int, message string) error {
	return o.JSON(code, lib.InRow{
		"code":     0,
		"success":  false,
		"msg":      "",
		"err":      message,
		"sub_code": "fail",
	})
}

//GZIPJSONFail
func (o *Output) GZIPJSONFail(code int, message string) error {
	return o.GZIPJSON(code, lib.InRow{
		"code":     0,
		"success":  false,
		"msg":      "",
		"err":      message,
		"sub_code": "fail",
	})
}

//JSONObject
func (o *Output) JSONObject(code int, info interface{}) error {
	return o.JSON(code, lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      "获取成功",
		"err":      "",
		"info":     info,
		"sub_code": "info.success",
	})
}


//GZIPJSONObject
func (o *Output) GZIPJSONObject(code int, info interface{}) error {
	return o.GZIPJSON(code, lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      "获取成功",
		"err":      "",
		"info":     info,
		"sub_code": "info.success",
	})
}


//JSONList
func (o *Output) JSONList(code int, list interface{}) error {
	return o.JSON(code, lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      "获取成功",
		"err":      "",
		"list":     list,
		"sub_code": "list.success",
	})
}


//GZIPJSONList
func (o *Output) GZIPJSONList(code int, list interface{}) error {
	return o.GZIPJSON(code, lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      "获取成功",
		"err":      "",
		"list":     list,
		"sub_code": "list.success",
	})
}


//Success 成功输出json
func (o *Output) SuccessJSON(data lib.InRow) error {
	return o.JSON(200, lib.MergeInRow(lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      "获取成功",
		"err":      "",
		"sub_code": "list.success",
	}, data))
}

//SuccessGZipJSON 成功输出json
func (o *Output) SuccessGZipJSON(data lib.InRow) error {
	return o.GZIPJSON(200, lib.MergeInRow(lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      "获取成功",
		"err":      "",
		"sub_code": "list.success",
	}, data))
}

//JSONPager
func (o *Output) JSONPager(code int, list interface{}, pager interface{}) error {
	return o.JSON(code, lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      "获取成功",
		"err":      "",
		"list":     list,
		"pager":    pager,
		"sub_code": "pager.success",
	})
}


//GZIPJSONPager
func (o *Output) GZIPJSONPager(code int, list interface{}, pager interface{}) error {
	return o.GZIPJSON(code, lib.InRow{
		"code":     1,
		"success":  true,
		"msg":      "获取成功",
		"err":      "",
		"list":     list,
		"pager":    pager,
		"sub_code": "pager.success",
	})
}

//JSONP
func (o *Output) JSONP(code int, value interface{}) error {
	return o.Render(code, value, binding.JSONP)
}

//Render
func (o *Output) Render(code int, value interface{}, bind binding.Rendering) error {
	if o.status == 0 && code > 0 {
		o.status = code
		//log.DebugPrint("writer:%v",o.writer)
		err := bind.Render(code, o.writer, value)
		return err
	}
	return nil
}

//HTML
func (o *Output) HTML(code int, name string, value interface{}) error {
	bind := binding.NewHTML(name, o.templ)
	return o.Render(code, value, bind)
}

//XML
func (o *Output) XML(code int, value interface{}) error {
	return o.Render(code, value, binding.XML)
}

//YAML
func (o *Output) YAML(code int, value interface{}) error {
	return o.Render(code, value, binding.YAML)
}

//Text
func (o *Output) Text(code int, value interface{}) error {
	return o.Render(code, value, binding.TEXT)
}

//Redirect 跳转
func (o *Output) Redirect(code int, location string) {
	if o.status == 0 && code > 0 {
		o.status = code
		http.Redirect(o.writer, o.in.R(), location, code)
	}
}

//NotFound 404
func (o *Output) NotFound() {
	if o.status == 0 {
		o.status = 404
		http.NotFound(o.writer, o.in.R())
	}
}

//HasOutput 是否输出了 true已经输出 false 未输出
func (o *Output) HasOutput() bool {
	if o.status == 0 {
		return false
	}
	return true
}

//NewInput 新建一个input
func NewOutput() *Output {
	return &Output{}
}
