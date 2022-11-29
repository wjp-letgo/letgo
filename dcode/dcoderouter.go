package dcode

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
)

//路由检测
type DCodeRouter struct {
	fset       *token.FileSet
	mfile      *ast.File
	modName    string
	controller []ControllerInfo
	comments   []*ast.CommentGroup
}

//Visit 遍历抽象语法树
func (c *DCodeRouter) Visit(node ast.Node) ast.Visitor {
	ck, ok := node.(*ast.CommentGroup)
	if ok {
		c.comments = append(c.comments, ck)
	}
	switch node.(type) {
	case *ast.File:
		c.astFile(node.(*ast.File))
		break
	case *ast.GenDecl:
		c.genDecl(node.(*ast.GenDecl))
		break
	case *ast.FuncDecl:
		c.funcDecl(node.(*ast.FuncDecl))
	}
	return c
}

//Finish 遍历完成
func (c *DCodeRouter) Finish() {
	//fmt.Println(c.funcList)
	c.mfile.Comments = c.comments
}

var routerNeedImport []string = []string{
	"github.com/wjp-letgo/letgo/web", "github.com/wjp-letgo/letgo/web/context", "github.com/wjp-letgo/letgo/web/filter",
}

//astFile 根
func (c *DCodeRouter) astFile(file *ast.File) {
	//ast.Print(c.fset,file.Imports)
	if _, ok := file.Scope.Objects["InitRouter"]; !ok {
		c.createFunc(file)
	}
	for _, v := range c.controller {
		//导入包
		//routerNeedImport=append(routerNeedImport,v.Path)
		pg := fmt.Sprintf("%s/%s", c.modName, v.Path)
		routerNeedImport = append(routerNeedImport, pg)
	}
	if file.Imports == nil {
		CreateImport(file, routerNeedImport...)
		//ast.Print(c.fset,file)
	} else {
		for _, s := range routerNeedImport {
			var find bool
			for _, v := range file.Imports {
				if v.Path.Value == strconv.Quote(s) {
					find = true
				}
			}
			if !find {
				CreateImport(file, s)
			}
		}
		//ast.Print(c.fset, file.Imports)
	}
}

//genDecl 字节点
func (c *DCodeRouter) genDecl(decl *ast.GenDecl) {
	//ast.Print(c.fset,decl)
}

//funcDecl 函数节点
func (c *DCodeRouter) funcDecl(fn *ast.FuncDecl) {
	//ast.Print(c.fset,fn)
	if fn.Name.Name == "InitRouter" {
		//在初始化路由里做手脚
		c.initRouter(fn)
	}
}

//创建run函数
func (c *DCodeRouter) createFunc(file *ast.File) {
	obj := &ast.Object{
		Name: "InitRouter",
		Kind: ast.Fun,
		Decl: &ast.FuncDecl{
			Type: &ast.FuncType{},
			Body: &ast.BlockStmt{},
		},
	}
	fn := &ast.FuncDecl{
		Name: &ast.Ident{
			Name: "InitRouter",
			Obj:  obj,
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
		},
		Body: &ast.BlockStmt{},
	}
	file.Decls = append(file.Decls, fn)
	file.Scope.Objects["InitRouter"] = obj
}

//initRouter 初始化路由添加控制器
func (c *DCodeRouter) initRouter(fn *ast.FuncDecl) {
	var filterFun bool
	ast.Inspect(fn.Body, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.SelectorExpr:
			n := node.(*ast.SelectorExpr)
			if n.Sel.Name == "AddFilter" {
				if x, ok := n.X.(*ast.Ident); ok && x.Name == "web" {
					filterFun = true
				}
			}
			break
		}
		return true
	})
	//ast.Print(c.fset, fn.Body)
	if !filterFun {
		c.createFilterFun(fn)
	}
	c.regiterRouter(fn)
	//ast.Print(c.fset,fn.Body)
}

//regiterRouter 注册路由
func (c *DCodeRouter) regiterRouter(fn *ast.FuncDecl) {
	//ast.Print(c.fset,fn.Body)
	//var x *ast.CallExpr
	var fns []ast.Stmt
	for _, v := range fn.Body.List {
		if cll, ok := v.(*ast.ExprStmt); ok {
			if cll2, ok2 := cll.X.(*ast.CallExpr); ok2 {
				if fun, ok3 := cll2.Fun.(*ast.SelectorExpr); ok3 {
					if fun.Sel.Name != "RegisterController" {
						fns = append(fns, v)
					}
				} else {
					fns = append(fns, v)
				}
			} else {
				fns = append(fns, v)
			}
		} else {
			fns = append(fns, v)
		}
	}
	fn.Body.List = fns
	//ast.Print(c.fset,fn.Body)
	for _, nc := range c.controller {
		c.createController(fn, nc)
	}
	//ast.Print(c.fset,fn.Body)
	/*
		if x == nil {
			//创建
			c.createController(fn, cs)
		} else {
			//更新
			c.updateController(x, cs)
		}
	*/
}

//createController 创建控制器
func (c *DCodeRouter) createController(fn *ast.FuncDecl, cs ControllerInfo) {
	//fmt.Println("创建控制器")
	var args []ast.Expr
	args = append(args, &ast.UnaryExpr{
		Op: token.AND,
		X: &ast.CompositeLit{
			Type: &ast.SelectorExpr{
				X: &ast.Ident{
					Name: cs.PackageName,
				},
				Sel: &ast.Ident{
					Name: cs.ControllerName,
				},
			},
			Incomplete: false,
		},
	})
	for _, fs := range cs.Funcs {
		args = append(args, &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(fmt.Sprintf("%s:%s", fs.Method, fs.MethodName)),
		})
	}
	fn.Body.List = append(fn.Body.List, &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: &ast.Ident{
					Name: "web",
				},
				Sel: &ast.Ident{
					Name: "RegisterController",
				},
			},
			Args: args,
		},
	})
}

//updateController 更新控制器
func (c *DCodeRouter) updateController(x *ast.CallExpr, cs ControllerInfo) {
	var ags []ast.Expr
	ags = append(ags, x.Args[0])
	for _, fs := range cs.Funcs {
		var find bool
		fnn := fmt.Sprintf("%s:%s", fs.Method, fs.MethodName)
		for _, v := range x.Args {
			if n, ok := v.(*ast.BasicLit); ok {
				if n.Value == strconv.Quote(fnn) {
					find = true
					ags = append(ags, v)
				}
			}
		}
		if !find {
			ags = append(ags, &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(fmt.Sprintf("%s:%s", fs.Method, fs.MethodName)),
			})
		}
	}
	//ast.Print(c.fset,ags)
	x.Args = ags
}

//createFilterFun 创建过滤函数
func (c *DCodeRouter) createFilterFun(fn *ast.FuncDecl) {
	fun := &ast.SelectorExpr{
		X: &ast.Ident{
			Name: "web",
		},
		Sel: &ast.Ident{
			Name: "AddFilter",
		},
	}
	var args []ast.Expr
	args = append(args, &ast.BasicLit{
		Kind:  token.STRING,
		Value: strconv.Quote("/*"),
	})
	args = append(args, &ast.SelectorExpr{
		X: &ast.Ident{
			Name: "filter",
		},
		Sel: &ast.Ident{
			Name: "BEFORE_ROUTER",
		},
	})
	var names []*ast.Ident
	names = append(names, &ast.Ident{
		Name: "ctx",
		Obj: &ast.Object{
			Kind: ast.Var,
			Name: "ctx",
			Decl: &ast.FuncDecl{
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{},
			},
		},
	})
	var args2 []*ast.Field
	args2 = append(args2, &ast.Field{
		Names: names,
		Type: &ast.StarExpr{
			X: &ast.SelectorExpr{
				X: &ast.Ident{
					Name: "context",
				},
				Sel: &ast.Ident{
					Name: "Context",
				},
			},
		},
	})
	args = append(args, &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: args2,
			},
		},
		Body: &ast.BlockStmt{},
	})
	x := &ast.CallExpr{
		Fun:  fun,
		Args: args,
	}
	bodyL := &ast.ExprStmt{
		X: x,
	}
	fn.Body.List = append(fn.Body.List, bodyL)
}
