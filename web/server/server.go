package server

import (
	syscontext "context"
	"net/http"
	"sync"

	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/log"
	"github.com/wjp-letgo/letgo/web/context"
	"github.com/wjp-letgo/letgo/web/router"
	"github.com/wjp-letgo/letgo/web/tmpl"
)

//Server httpServer类
type Server struct {
	pool sync.Pool
	route *router.Router
	httpServer *http.Server
}
//NewServer
func NewServer()*Server{
	sr:= &Server{
		route: router.HttpRouter(),
	}
	sr.pool.New=func()interface{}{
		//log.DebugPrint("新建一个context")
		return context.NewContext()
	}
	return sr
}
//ServeHTTP http请求
func (s *Server) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	c:=s.pool.Get().(*context.Context)
	c.Request=r
	c.Writer=w
	//log.DebugPrint("请求过来了:%p,%p",&w,&c.Writer)
	c.Reset()
	s.handleHttpRequest(c)
	s.pool.Put(c)
}
//handleHttpRequest 处理http请求
func (s *Server)handleHttpRequest(c *context.Context){
	s.route.HandleHttpRequest(c)
}
//Run 启动服务
func (s *Server)Run(addr ...string)error {
	address:=lib.ResolveAddress(addr)
	//fmt.Println(address)
	log.DebugPrint("Start server address:%s",address)
	s.httpServer = &http.Server{Addr: address, Handler: s}
	return s.httpServer.ListenAndServe()
}

//RunTLS 启动服务
func(s *Server)RunTLS(certFile, keyFile string, addr ...string)error{
	address:=lib.ResolveAddress(addr)
	s.httpServer = &http.Server{Addr: address, Handler: s}
	return s.httpServer.ListenAndServeTLS(certFile, keyFile)
}
//RegisterRouter 注册路由
func(s *Server)RegisterRouter(method,relativePath string, handler context.HandlerFunc){
	s.route.RegisterRouter(method,relativePath,handler)
}

//Router 获得路由
func (s *Server)Router()*router.Router{
	return s.route
}
//Tmpl 获得模板对象
func (s *Server)Tmpl()*tmpl.Tmpl{
	return tmpl.GetTmpl()
}
//Shutdown 关闭服务
func (s *Server)Shutdown(ctx syscontext.Context)error{
	return s.httpServer.Shutdown(ctx)
}