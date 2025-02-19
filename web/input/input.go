package input

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/wjp-letgo/letgo/file"
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/log"
	"github.com/wjp-letgo/letgo/web/binding"
	"github.com/wjp-letgo/letgo/web/headerlock"
)

const defaultMultipartMem = 64 << 20 //64MB

type UploadFunc func(*multipart.FileHeader) //上传回调

// Input
type Input struct {
	Method  string
	params  lib.Row
	get     lib.Row
	post    lib.Row
	request *http.Request
	body    []byte
}

// R 获取请求
func (i *Input) R() *http.Request {
	return i.request
}

// Param 获得参数
func (i *Input) Param(key string, value ...interface{}) *lib.Data {
	if _, ok := i.params[key]; !ok {
		if len(value) > 0 && value[0] != nil {
			return &lib.Data{Value: value[0]}
		}
	}
	if i.params[key] == nil && len(value) > 0 && value[0] != nil {
		return &lib.Data{Value: value[0]}
	}
	return i.params[key]
}

// SetParam 设置参数
func (i *Input) SetParam(key string, value interface{}) {
	i.set("params", key, value)
}

// Param 获得参数
func (i *Input) ParamRaw() lib.Row {
	return i.params
}

// Get 获得参数
func (i *Input) Get(key string, value ...interface{}) *lib.Data {
	if _, ok := i.get[key]; !ok {
		if len(value) > 0 && value[0] != nil {
			return &lib.Data{Value: value[0]}
		}
	}
	if i.get[key] == nil && len(value) > 0 && value[0] != nil {
		return &lib.Data{Value: value[0]}
	}
	return i.get[key]
}

// GetRaw 获得参数
func (i *Input) GetRaw() lib.Row {
	return i.get
}

// SetGet 设置参数
func (i *Input) SetGet(key string, value interface{}) {
	i.set("get", key, value)
}

// BodyBytes 内容
func (i *Input) BodyBytes() []byte {
	return i.body
}

// Body
func (i *Input) Body() string {
	b := i.BodyBytes()
	if b != nil {
		return string(b)
	}
	return ""
}

// Post 获得参数
func (i *Input) Post(key string, value ...interface{}) *lib.Data {
	if _, ok := i.post[key]; !ok {
		if len(value) > 0 && value[0] != nil {
			return &lib.Data{Value: value[0]}
		}
	}
	if i.post[key] == nil && len(value) > 0 && value[0] != nil {
		return &lib.Data{Value: value[0]}
	}
	return i.post[key]
}

// GetRaw 获得参数
func (i *Input) PostRaw() lib.Row {
	return i.post
}

// SetPost 设置参数
func (i *Input) SetPost(key string, value interface{}) {
	i.set("post", key, value)
}

// set 设置参数
func (i *Input) set(method, key string, value interface{}) {
	switch method {
	case "params":
		i.params[key] = &lib.Data{Value: value}
	case "get":
		i.get[key] = &lib.Data{Value: value}
	case "post":
		i.post[key] = &lib.Data{Value: value}
	}
}

// Init 初始化
func (i *Input) Init(request *http.Request) {
	i.Method = request.Method
	i.request = request
	if request.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(i.request.Body)
		if err != nil {
			log.DebugPrint("error on parse body gzip: %v", err)
			return
		}
		defer gzipReader.Close()
		i.body, _ = ioutil.ReadAll(gzipReader)
	} else {
		i.body, _ = ioutil.ReadAll(i.request.Body)
	}
	i.request.Body.Close()
	i.request.Body = ioutil.NopCloser(bytes.NewBuffer(i.body))
	if !strings.Contains(request.Header.Get("content-type"), "application/json") {
		err := i.request.ParseMultipartForm(defaultMultipartMem)
		if err != nil {
			if err != http.ErrNotMultipart {
				log.DebugPrint("error on parse multipart form array: %v", err)
			}
		}
	}
	i.initQuery()
	i.initForm()
	i.initJson()
}

// initJson
func (i *Input) initJson() {
	if strings.Index(i.ContentType(), "application/json") != -1 {
		var result map[string]interface{}
		lib.StringToObject(i.Body(), &result)
		for k, v := range result {
			i.SetParam(k, v)
			i.SetPost(k, v)
		}
	}
}

// initQuery
func (i *Input) initQuery() {
	query := i.request.URL.Query()
	for k, v := range query {
		if len(v) == 1 {
			i.SetParam(k, v[0])
			i.SetGet(k, v[0])
		} else {
			i.SetParam(k, v)
			i.SetGet(k, v)
		}
	}
}

// File
func (i *Input) File(name string) (*multipart.FileHeader, error) {
	if i.request.MultipartForm == nil {
		return nil, errors.New("MultipartForm is nil")
	}
	if fhs := i.request.MultipartForm.File[name]; len(fhs) > 0 {
		f, err := fhs[0].Open()
		if err != nil {
			return nil, err
		}
		f.Close()
		return fhs[0], nil
	}
	return nil, errors.New("file does not exist")
}

// SaveUploadFile
func (i *Input) SaveUploadFile(name, dst string) error {
	f, err := i.File(name)
	if err != nil {
		return err
	}
	src, err := f.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	file.Mkdir(file.DirName(dst))
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}

// SaveUploadFileByDir
func (i *Input) SaveUploadFileByDir(name, dir string) (string, error) {
	f, err := i.File(name)
	if err != nil {
		return "", err
	}
	src, err := f.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	file.Mkdir(file.DirName(dir))
	if dir[len(dir)-1:] == "/" {
		dir = dir + f.Filename
	} else {
		dir = dir + "/" + f.Filename
	}
	out, err := os.Create(dir)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return dir, err
}

// SaveUploadByFunc
func (i *Input) SaveUploadByFunc(name string, callback UploadFunc) error {
	f, err := i.File(name)
	if err != nil {
		return err
	}
	callback(f)
	return nil
}

// initForm
func (i *Input) initForm() {
	for k, v := range i.request.PostForm {
		if len(v) == 1 {
			i.SetParam(k, v[0])
			i.SetPost(k, v[0])
		} else {
			i.SetParam(k, v)
			i.SetPost(k, v)
		}
	}
}

// ClientIp 客户端IP
func (i *Input) ClientIp() string {
	var addr string
	headerlock.HeaderMapMutex.RLock()
	addr = i.request.Header.Get("X-Appengine-Remote-Addr")
	headerlock.HeaderMapMutex.RUnlock()
	if addr != "" {
		return addr
	}
	headerlock.HeaderMapMutex.RLock()
	addr = i.request.Header.Get("X-Forwarded-For")
	headerlock.HeaderMapMutex.RUnlock()
	if addr != "" {
		return addr
	}
	headerlock.HeaderMapMutex.RLock()
	addr = i.request.Header.Get("X-real-ip")
	headerlock.HeaderMapMutex.RUnlock()
	if addr != "" {
		return addr
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(i.request.RemoteAddr))
	if err != nil {
		return ""
	}
	return ip
}

// ContentType
func (i *Input) ContentType() string {
	headerlock.HeaderMapMutex.RLock()
	ct := i.request.Header.Get("Content-Type")
	headerlock.HeaderMapMutex.RUnlock()
	return ct

}

// Bind
func (i *Input) Bind(value interface{}) error {
	bind := binding.NewBind(i.ContentType())
	return i.BindWith(value, bind)
}

// BindJSON
func (i *Input) BindJSON(value interface{}) error {
	return i.BindWith(value, binding.JSON)
}

// BindJSONByKey
func (i *Input) BindJSONByKey(key string, value interface{}) bool {
	data := i.Param(key).Value
	return lib.StringToObject(lib.ObjectToString(data), value)
}

// BindXML
func (i *Input) BindXML(value interface{}) error {
	return i.BindWith(value, binding.XML)
}

// BindYAML
func (i *Input) BindYAML(value interface{}) error {
	return i.BindWith(value, binding.YAML)
}

// BindWith
func (i *Input) BindWith(value interface{}, bind binding.Binding) error {
	return bind.Bind(i.request, i.body, value)
}

// NewInput 新建一个input
func NewInput() *Input {
	return &Input{
		params: make(lib.Row),
		get:    make(lib.Row),
		post:   make(lib.Row),
	}
}
