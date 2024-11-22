package httpclient

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"strings"
	"time"

	"github.com/wjp-letgo/letgo/file"
	"github.com/wjp-letgo/letgo/lib"
)

const (
	OPT_SSL int =iota
	OPT_PROXY
	PROXY_SOCKS4
	PROXY_SOCKS5
	PROXY_SOCKS4A
	PROXY_HTTP
	PROXY_HTTPS
	OPT_TIMEOUT
)
var protocol=map[int]string{
	PROXY_SOCKS4:"socks4",
	PROXY_SOCKS5:"socks5",
	PROXY_SOCKS4A:"socks4a",
	PROXY_HTTP:"http",
	PROXY_HTTPS:"https",
}
var constVar=lib.IntRow{
	OPT_SSL:"OPT_SSL",
	OPT_PROXY:"OPT_PROXY",
	OPT_TIMEOUT:"OPT_TIMEOUT",
}
//HttpResponse 响应内容
type HttpResponse struct {
	Code int `json:"code"`
	Err string `json:"err"`
	BodyByte []byte `json:"bodybyte"`
	Header http.Header
	Dump string `json:"dump"`
}
//RequestBeforeFunc 请求前函数
type RequestBeforeFunc func(*http.Request)

//String 字符输出
func (h *HttpResponse)String()string{
	return lib.ObjectToString(h)
}
//Body 内容
func (h *HttpResponse)Body()string{
	return string(h.BodyByte)
}
//Httper http请求
type Httper interface{
	WithHeader(key string, value interface{}) Httper
	WithHeaders(header lib.InRow) Httper
	WithOption(key int, value interface{}) Httper
	WithOptions(options lib.IntRow) Httper
	WithCookie(cookies ...*http.Cookie) Httper
	WithSSL(certpemPath, keypemPath, rootcaName string) Httper
	WithProxy(proto int, host, port string) Httper
	WithTimeOut(timeout int) Httper
	Get(url string,values lib.InRow)*HttpResponse
	Put(url string,values lib.InRow) *HttpResponse
	Post(url string,values lib.InRow)*HttpResponse
	PostJson(url string,value interface{})*HttpResponse
	PostXml(url string,value interface{})*HttpResponse
	PostMultipart(url string,values lib.InRow)*HttpResponse
	Delete(url string,values lib.InRow)*HttpResponse
	Options(url string,values lib.InRow)*HttpResponse
	Head(url string)*HttpResponse
	Connect(url string,values lib.InRow)*HttpResponse
	Trace(url string, values lib.InRow)*HttpResponse
	Patch(url string,values lib.InRow)*HttpResponse
	WithRequestBefore(fun RequestBeforeFunc)Httper
}
//Http 类
type HttpClient struct {
	options lib.IntRow
	headers lib.InRow
	cookie []*http.Cookie
	fun RequestBeforeFunc
	dump string
}
//WithHeader 自定义头
func(h *HttpClient)WithHeader(key string, value interface{})Httper{
	if h.headers==nil{
		h.headers=make(lib.InRow)
	}
	h.headers[key]=value
	return h
}
//WithHeaders 自定义头
func(h *HttpClient)WithHeaders(header lib.InRow) Httper{
	for k,v:=range header{
		h.WithHeader(k,v)
	}
	return h
}
//WithRequestBefore 请求前调用
func(h *HttpClient)WithRequestBefore(fun RequestBeforeFunc)Httper{
	h.fun=fun
	return h
}
//WithOption 自定义选项
func(h *HttpClient)WithOption(key int, value interface{}) Httper{
	if _,ok:=constVar[key];!ok{
		panic("Option does not exist");
	}
	if h.options==nil{
		h.options=make(lib.IntRow)
	}
	h.options[key]=value
	return h
}
//WithOption 自定义选项
func(h *HttpClient)WithOptions(options lib.IntRow) Httper{
	for k,v:=range options{
		h.WithOption(k,v)
	}
	return h
}
//WithCookie 自定义cookie
func(h *HttpClient)WithCookie(cookies ...*http.Cookie) Httper{
	h.cookie=append(h.cookie, cookies...)
	return h
}
//WithSSL 设置证书
func(h *HttpClient)WithSSL(certpemPath, keypemPath, rootcaName string) Httper{
	h.WithOption(OPT_SSL, lib.InRow{
		"certpemPath":certpemPath,
		"keypemPath":keypemPath,
		"rootcaName":rootcaName,
	})
	return h
}
//WithProxy 代理
func(h *HttpClient)WithProxy(proto int, host, port string) Httper{
	h.WithOption(OPT_PROXY, lib.InRow{
		"proto":proto,
		"host":host,
		"port":port,
	})
	return h
}
//WithTimeOut 超时时间 单位秒
func(h *HttpClient)WithTimeOut(timeout int) Httper{
	h.WithOption(OPT_TIMEOUT, timeout)
	return h
}
//Get 请求
func(h *HttpClient)Get(url string,values lib.InRow) *HttpResponse{
	if len(values)>0{
		if !strings.Contains(url,"?") {
			url=url+"?"+HttpBuildQuery(values)
		}else{
			if !strings.Contains(url,"&") {
				url=url+HttpBuildQuery(values)
			} else{
				url=url+"&"+HttpBuildQuery(values)
			}
			
		}
	}
	client:=h.getClient()
	req:=h.getRequest(url,"GET",nil)
	return h.getResponse(client.Do(req))
}
//Post 请求
func(h *HttpClient)Post(url string,values lib.InRow) *HttpResponse{
	if h.checkIncludeFile(values) {
		return h.PostMultipart(url,values)
	}
	if _,ok:=h.headers["Content-Type"];!ok{
		h.WithHeader("Content-Type", "application/x-www-form-urlencoded")
	}
	body:=strings.NewReader(HttpBuildQuery(values))
	client:=h.getClient()
	req:=h.getRequest(url,"POST",body)
	return h.getResponse(client.Do(req))
}
//PostJson 请求json
func(h *HttpClient)PostJson(url string,value interface{}) *HttpResponse{
	if _,ok:=h.headers["Content-Type"];!ok{
		h.WithHeader("Content-Type", "application/json")
	}
	var body []byte
	switch t := value.(type) {
	case []byte:
		body = t
	case string:
		body = []byte(t)
	default:
		var err error
		body, err = json.Marshal(value)
		if err != nil {
			return h.responseErr(err)
		}
	}
	//fmt.Println("body",string(body))
	client:=h.getClient()
	req:=h.getRequest(url,"POST",bytes.NewReader(body))
	return h.getResponse(client.Do(req))
}
//PostXml 请求xml
func(h *HttpClient)PostXml(url string,value interface{})*HttpResponse {
	if _,ok:=h.headers["Content-Type"];!ok{
		h.WithHeader("Content-Type", "text/xml")
	}
	var body []byte
	switch t := value.(type) {
	case []byte:
		body = t
	case string:
		body = []byte(t)
	default:
		var err error
		body, err = xml.Marshal(value)
		if err != nil {
			return h.responseErr(err)
		}
	}
	client:=h.getClient()
	req:=h.getRequest(url,"POST",bytes.NewReader(body))
	return h.getResponse(client.Do(req))
}
//PostMultipart 请求提交文件
func(h *HttpClient)PostMultipart(url string,values lib.InRow)*HttpResponse{
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k,v:=range values{
		if k[0] == '@' {
			var err error
			if vs, ok := v.(string); ok {
				err=h.addFile(writer,k[1:],vs)
			}else if vs, ok := v.([]byte); ok {
				err=h.addFileBytes(writer,k[1:],vs)
			}
			if err!=nil{
				return h.responseErr(err)
			}
		}else{
			writer.WriteField(k,(&lib.Data{Value:v}).String())
		}
	}
	writer.Close()
	if _,ok:=h.headers["Content-Type"];!ok{
		h.WithHeader("Content-Type",writer.FormDataContentType())
	}
	h.WithHeader("Content-Length",len(body.Bytes()))
	client:=h.getClient()
	req:=h.getRequest(url,"POST",body)
	return h.getResponse(client.Do(req))
}
//addFile 添加文件
func(h *HttpClient)addFile(writer *multipart.Writer, name,path string)error{
	part,err:=writer.CreateFormFile(name,file.BaseName(path))
	if err!=nil{
		return err
	}
	_,err=part.Write([]byte(file.GetContent(path)))
	return err
}
//addFileBytes
func(h *HttpClient)addFileBytes(writer *multipart.Writer, name string,fileData []byte)error{
	k := make(textproto.MIMEHeader)
	k.Set("Content-Disposition",fmt.Sprintf(`form-data; name="%s"; filename="%s"`,name,name))
	k.Set("Content-Type", "application/octet-stream")
	part,err:=writer.CreatePart(k)
	if err!=nil{
		return err
	}
	_,err=part.Write(fileData)
	return err
}
//Delete 请求
func(h *HttpClient)Delete(url string,values lib.InRow) *HttpResponse{
	if len(values)>0{
		if !strings.Contains(url,"?") {
			url=url+"?"+HttpBuildQuery(values)
		}else{
			url=url+HttpBuildQuery(values)
		}
	}
	client:=h.getClient()
	req:=h.getRequest(url,"DELETE",nil)
	return h.getResponse(client.Do(req))
}
//Options 请求
func(h *HttpClient)Options(url string,values lib.InRow) *HttpResponse{
	if len(values)>0{
		if !strings.Contains(url,"?") {
			url=url+"?"+HttpBuildQuery(values)
		}else{
			url=url+HttpBuildQuery(values)
		}
	}
	client:=h.getClient()
	req:=h.getRequest(url,"OPTIONS",nil)
	return h.getResponse(client.Do(req))
}
//Head 请求
func(h *HttpClient)Head(url string) *HttpResponse{
	client:=h.getClient()
	req:=h.getRequest(url,"HEAD",nil)
	return h.getResponse(client.Do(req))
}

//Put 请求
func(h *HttpClient)Put(url string,values lib.InRow) *HttpResponse{
	if h.checkIncludeFile(values) {
		return h.PostMultipart(url,values)
	}
	if _,ok:=h.headers["Content-Type"];!ok{
		h.WithHeader("Content-Type", "application/x-www-form-urlencoded")
	}
	body:=strings.NewReader(HttpBuildQuery(values))
	client:=h.getClient()
	req:=h.getRequest(url,"PUT",body)
	return h.getResponse(client.Do(req))
}

//Connect 请求
func(h *HttpClient)Connect(url string,values lib.InRow) *HttpResponse{
	if len(values)>0{
		if !strings.Contains(url,"?") {
			url=url+"?"+HttpBuildQuery(values)
		}else{
			url=url+HttpBuildQuery(values)
		}
	}
	client:=h.getClient()
	req:=h.getRequest(url,"CONNECT",nil)
	return h.getResponse(client.Do(req))
}

//Trace 请求
func(h *HttpClient)Trace(url string,values lib.InRow) *HttpResponse{
	if len(values)>0{
		if !strings.Contains(url,"?") {
			url=url+"?"+HttpBuildQuery(values)
		}else{
			url=url+HttpBuildQuery(values)
		}
	}
	client:=h.getClient()
	req:=h.getRequest(url,"TRACE",nil)
	return h.getResponse(client.Do(req))
}
//Patch 请求
func(h *HttpClient)Patch(url string,values lib.InRow) *HttpResponse{
	if len(values)>0{
		if !strings.Contains(url,"?") {
			url=url+"?"+HttpBuildQuery(values)
		}else{
			url=url+HttpBuildQuery(values)
		}
	}
	client:=h.getClient()
	req:=h.getRequest(url,"PATCH",nil)
	return h.getResponse(client.Do(req))
}

//getClient 获得客户端
func (h *HttpClient)getClient()*http.Client{
	transport:=h.getTransport()
	client:=&http.Client{
		Transport: transport,
	}
	return client
}
//checkIncludeFile 是否包含上传文件
func (h *HttpClient)checkIncludeFile(values lib.InRow)bool{
	for k, _:= range values {
		if k[0] == '@' {
			return true
		}
	}
	return false
}
//getTransport 获得getTransport
func (h *HttpClient)getTransport()*http.Transport{
	transport:=&http.Transport{}
	if len(h.options)==0{
		transport.TLSClientConfig=&tls.Config{InsecureSkipVerify: true}
		return transport	
	}
	for optionKey,option:=range h.options{
		if optionKey==OPT_SSL {
			ssl:=option.(lib.InRow)
			cert, err :=tls.LoadX509KeyPair(ssl["certpemPath"].(string),ssl["keypemPath"].(string))
			if err!=nil{
				transport.TLSClientConfig=&tls.Config{InsecureSkipVerify: true}
			}else{
				pool := x509.NewCertPool()
				content:=file.GetContent(ssl["keypemPath"].(string))
				pool.AppendCertsFromPEM([]byte(content))
				transport.TLSClientConfig = &tls.Config{RootCAs: pool, Certificates: []tls.Certificate{cert}}
			}
		} else if optionKey==OPT_PROXY{
			proxy:=option.(lib.InRow)
			address:=fmt.Sprintf("%s://%s:%s",protocol[proxy["proto"].(int)],proxy["host"].(string),proxy["port"].(string))
			//address:=fmt.Sprintf("%s:%s",proxy["host"].(string),proxy["port"].(string))
			transport.Proxy=func(_ *http.Request)(*url.URL, error){
				return url.Parse(address)
			}
		} else if optionKey==OPT_TIMEOUT{
			connectTimeout:=option.(int)
			transport.Dial=func(network, addr string) (net.Conn, error) {
				var conn net.Conn
				var err error
				if connectTimeout > 0 {
					conn, err = net.DialTimeout(network, addr, time.Duration(connectTimeout)*time.Second)
					if err != nil {
						return nil, err
					}
				} else {
					conn, err = net.Dial(network, addr)
					if err != nil {
						return nil, err
					}
				}
				if connectTimeout > 0 {
					conn.SetDeadline(time.Now().Add(time.Duration(connectTimeout)*time.Second))
				}
				return conn, nil
			}
		}
	}
	return transport
}
//getRequest 发起请求
func (h *HttpClient)getRequest(url,method string,body io.Reader)*http.Request{
	req,err:=http.NewRequest(method, url, body)
	if err!=nil{
		return nil
	}
	for headerKey,headerValue:=range h.headers {
		req.Header.Set(headerKey,(&lib.Data{Value: headerValue}).String())
	}
	for _,cookie:=range h.cookie{
		req.AddCookie(cookie)
	}
	dump, _ := httputil.DumpRequestOut(req, true)
	h.dump=fmt.Sprintf("%s", dump)
	if (h.fun!=nil){
		h.fun(req)
	}
	return req
}
//responseErr 返回错误
func (h *HttpClient)responseErr(err error)*HttpResponse{
	result:=&HttpResponse{}
	result.BodyByte=nil
	result.Err=err.Error()
	result.Code=500
	return result
}
//getResponse 获得响应内容
func (h *HttpClient)getResponse(response *http.Response,err error) *HttpResponse{
	if err!=nil {
		return h.responseErr(err)
	}
	dumprs,_:=httputil.DumpResponse(response,true)
	result:=&HttpResponse{}
	defer response.Body.Close()
	var reader io.ReadCloser
	encode:=response.Header.Get("Content-Encoding")
	if encode=="gzip" {
		reader,err=gzip.NewReader(response.Body)
		if err!=nil {
			return h.responseErr(err)
		}
	}else{
		reader=response.Body
	}
	defer reader.Close()
	content,err:=ioutil.ReadAll(reader)
	if err!=nil {
		return h.responseErr(err)
	}
	result.BodyByte=content
	result.Code=response.StatusCode
	result.Err=""
	result.Header=response.Header
	result.Dump=h.dump+fmt.Sprintf("\n\n===================================\n\n%s",dumprs) 
	return result
}
//HttpBuildQuery 生成 URL-encode 之后的请求字符串
func HttpBuildQuery(values lib.InRow) string{
	paramsValue:=make(url.Values)
	for k, v := range values {
		switch t:=v.(type) {
		case []string:
			for _,s:=range t{
				paramsValue.Add(k,s)
			}
		case []int64:
			for _,s:=range t{
				paramsValue.Add(k,fmt.Sprintf("%d",s))
			}
		default:
			vs:=fmt.Sprint(v)
			paramsValue.Add(k,vs)
		}
		
	}
	return paramsValue.Encode()
}

//UrlEncode url编码
func UrlEncode(queryString string)string{
	qs:=url.QueryEscape(queryString)
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(qs, "+", "%20"),"*", "%2A"),"%7E","~")
}

//UrlDecode url解码
func UrlDecode(queryString string)string{
	qs,_:=url.QueryUnescape(queryString)
	return qs
}

//GetUrlParam
func GetUrlParam(ul,key string)string{
	u,_:=url.Parse(ul)
	m, _ := url.ParseQuery(u.RawQuery)
	if _,ok:=m[key];ok{
		return m[key][0]
	}
	return ""
}


//New
func New()*HttpClient{
	return &HttpClient{}
}

//保存远程文件到本地
//remoteFile远程文件地址
//localFullName本地文件地址
func SaveRemoteFile(remoteFile, localFullName string) bool {
	if file.FileExist(localFullName) {
		return true
	}
	ihttp:=New().WithTimeOut(120)
	rs:=ihttp.Get(remoteFile,nil)
	if rs.Code == 200 {
		file.PutContent(localFullName,rs.Body())
		return true
	}
	return false
}