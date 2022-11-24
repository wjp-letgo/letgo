package context

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/web/headerlock"
	"github.com/wjpxxx/letgo/web/input"
	"github.com/wjpxxx/letgo/web/output"
	"github.com/wjpxxx/letgo/web/session"
	"github.com/wjpxxx/letgo/web/tmpl"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

//HandlerFunc 请求处理函数
type HandlerFunc func(*Context)

//Context 上下文对象
type Context struct {
	Request    *http.Request
	Writer     http.ResponseWriter
	Input      *input.Input
	Output     *output.Output
	Session    session.Sessioner
	RouterPath string
	sameSite   http.SameSite
	Tmpl       *tmpl.Tmpl
}

//重置
func (c *Context) Reset() {
	c.Input = input.NewInput()
	c.Output = output.NewOutput()
	c.Session = session.GetSession(c)
	c.RouterPath = ""
}

//Init 初始化
func (c *Context) Init() {
	c.Session.Start() //启动session
	c.Input.Init(c.Request)
	c.Output.Init(c.Writer, c.Input, c.Tmpl.Template)
}

//FullPath
func (c *Context) FullPath() string {
	return c.Request.URL.String()
}

//Router
func (c *Context) Router() string {
	requestPath := strings.ToLower(c.Request.URL.Path)
	if requestPath == "" {
		return "/"
	}
	return requestPath
}

//SetCookie 设置cookie
func (c *Context) SetCookies(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	if path == "" {
		path = "/"
	}
	cookie := http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: c.sameSite,
		Secure:   secure,
		HttpOnly: httpOnly,
	}
	//log.DebugPrint("设置cookie:%p",&c.Writer)
	headerlock.HeaderMapMutex.Lock()
	http.SetCookie(c.Writer, &cookie)
	headerlock.HeaderMapMutex.Unlock()
}

//SetCookie 设置cookie
func (c *Context) SetCookie(name, value string) {
	c.SetCookies(name, value, 3600, "", c.Domain(), false, false)
}

//SetCookieByExpire
func (c *Context) SetCookieByExpire(name, value string, expire int) {
	//log.DebugPrint("host:%s,name:%s,value:%s",c.Host(),name,value)
	c.SetCookies(name, value, expire, "", c.Domain(), false, false)
}

//Host
func (c *Context) Host() string {
	hostArray := strings.Split(c.Request.Host, ":")
	return hostArray[0]
}

//HttpOrigin
func (c *Context) HttpOrigin() string {
	headerlock.HeaderMapMutex.RLock()
	r := c.Request.Header.Get("Origin")
	headerlock.HeaderMapMutex.RUnlock()
	return r
}

//Domain
func (c *Context) Domain() string {
	domain1 := lib.GetRootDomain(c.HttpOrigin())
	if domain1 != "" && domain1 != c.Host() {
		return domain1
	}
	return c.Host()
}

//Cookie 获得cookie
func (c *Context) Cookie(name string) *lib.Data {
	headerlock.HeaderMapMutex.RLock()
	cookie, err := c.Request.Cookie(name)
	headerlock.HeaderMapMutex.RUnlock()
	if err != nil {
		return nil
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return &(lib.Data{Value: val})
}

//SetSameSite
func (c *Context) SetSameSite(sameSite http.SameSite) {
	c.sameSite = sameSite
}

//ContentType
func (c *Context) ContentType() string {
	return c.Input.ContentType()
}

//DumpRequest
func (c *Context) DumpRequest() string {
	dump, _ := httputil.DumpRequest(c.Request, true)
	return string(dump)
}

//GetHeader
func (c *Context) GetHeader(key string) string {
	headerlock.HeaderMapMutex.RLock()
	r := c.Request.Header.Get(key)
	headerlock.HeaderMapMutex.RUnlock()
	return r
}

//NewContext 新建一个上下文
func NewContext() *Context {
	ctx := &Context{
		Input:  input.NewInput(),
		Output: output.NewOutput(),
		Tmpl:   tmpl.GetTmpl(),
	}
	//ctx.Session=session.GetSession(ctx)
	return ctx
}
