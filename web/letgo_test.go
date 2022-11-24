package web

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/web/limiting"
	"github.com/wjpxxx/letgo/web/context"
	"fmt"
	"testing"
	"github.com/wjpxxx/letgo/web/filter"
)


func TestLetgo(t *testing.T) {
	EnableLimiting(limiting.LIMIT_FLOW_TOKEN_BUCKET, 0.02)
	AddFilter("/user/*",filter.BEFORE_ROUTER,func(ctx *context.Context){
		ctx.Output.Redirect(302,"/")
	})
	Static("/assets/", "./assets")
	StaticFile("/1.png", "./assets/b3.jpg")
	LoadHTMLGlob("config/*")
	Get("/", func(ctx *context.Context){
		//ctx.Output.Redirect(301,"http://www.baidu.com")
		x:=ctx.Input.Param("a")
		if x!=nil{
			ctx.Session.Set("a",x.Int())
			var a int
			ctx.Session.Get("a",&a)
			fmt.Println("a:",a,"x:",x.Int())
		}
		ctx.Output.HTML(200,"index.tmpl",lib.InRow{
			"title":"wjp",
		})
	})
	Get("/text", func(ctx *context.Context){
		ctx.Output.Text(200,"Hello letgo")
	})
	
	Post("/user/:id([0-9]+)", func(ctx *context.Context){
		type A struct{
			Data string `json:"data"`
			ShopID int64 `json:"shop_id"`
		}
		a:=A{}
		ctx.SetCookie("a", "234234")
		ctx.Input.BindJSON(&a)
		fmt.Println(a)
		ctx.Output.YAML(500,lib.InRow{
			"message":"123123",
			"b":2,
		})
	})
	Get("/www/*",func(ctx *context.Context){
		ctx.Output.JSON(200,lib.InRow{
			"success":true,
		})
	})
	Get("/user/:id([0-9]+)/:id3([0-9]+)", func(ctx *context.Context){
		ctx.Output.XML(200,lib.InRow{
			"message":"123123",
			"b":2,
		})
	})
	c:=&UserController{}
	RegisterController(c)
	Run()
}
type UserController struct{}
func (c *UserController)Add(ctx *context.Context){
	ctx.Output.JSON(200,lib.InRow{
		"a":1,
		"b":"wjp",
	})
}