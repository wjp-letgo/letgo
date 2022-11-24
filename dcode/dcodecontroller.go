package dcode

import (
	"fmt"
	"github.com/wjpxxx/letgo/lib"
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

//控制器检测
type DCodeController struct {
	fset           *token.FileSet
	mfile          *ast.File
	comments       []*ast.CommentGroup
	funcList       []FuncList
	PackageName    string
	ControllerName string
}

//FuncList
type FuncList struct {
	MethodName string
	Method     string
	Controller string
}

var needImport map[string]*packageInfo

//packageInfo
type packageInfo struct {
	name         string
	need         bool
	fileHasUse   bool
	importHasUse bool
}

//getNeedImport
func getNeedImport() map[string]*packageInfo {
	ret := make(map[string]*packageInfo)
	ret["context"] = &packageInfo{name: "github.com/wjpxxx/letgo/web/context", need: false}
	ret["lib"] = &packageInfo{name: "github.com/wjpxxx/letgo/lib", need: false}
	ret["fmt"] = &packageInfo{name: "fmt", need: false}
	return ret
}

//Visit 遍历抽象语法树
func (c *DCodeController) Visit(node ast.Node) ast.Visitor {
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
func (c *DCodeController) Finish() {
	//fmt.Println(c.funcList)
	c.mfile.Comments = c.comments
	//fmt.Println(c.ControllerName)
}

//获得控制器的方法列表
func (c *DCodeController) GetFuncList() []FuncList {
	return c.funcList
}

//astFile 根
func (c *DCodeController) astFile(file *ast.File) {
	//ast.Print(c.fset, file.Unresolved)
	//ast.Print(c.fset, file.Scope.Objects)
	c.PackageName = file.Name.Name
	c.needContext(file)
	if file.Scope != nil {
		for k, _ := range file.Scope.Objects {
			if strings.Index(k, "Get") != 0 && strings.Index(k, "Controller") != -1 {
				c.ControllerName = k
			}
		}
	}
}

//genDecl 字节点
func (c *DCodeController) genDecl(decl *ast.GenDecl) {
	//ast.Print(c.fset,decl)
	if decl.Tok == token.IMPORT {
		//导入节点
		c.addImport(decl)
	}
}

//funcDecl 函数节点
func (c *DCodeController) funcDecl(fn *ast.FuncDecl) {
	c.genDoc(fn)
	//ast.Print(c.fset,fn)
}

//是否包含github.com/wjpxxx/letgo/web/context
func (c *DCodeController) needContext(file *ast.File) {
	for _, v := range file.Unresolved {
		for k2, _ := range needImport {
			if v.Name == k2 {
				needImport[k2].fileHasUse = true
			}
		}
	}
	for _, v := range file.Imports {
		for k2, v2 := range needImport {
			if v.Path.Value == strconv.Quote(v2.name) {
				needImport[k2].importHasUse = true
			}
		}
	}
	if file.Imports == nil {
		var importStrs []string
		for _, v2 := range needImport {
			if v2.fileHasUse && !v2.importHasUse {
				importStrs = append(importStrs, v2.name)
			}
		}
		if len(importStrs) > 0 {
			CreateImport(file, importStrs...)
		}
	} else {
		for k2, v2 := range needImport {
			if v2.fileHasUse && !v2.importHasUse {
				needImport[k2].need = true
			} else {
				needImport[k2].need = false
			}
		}
	}
	//ast.Print(c.fset, file)
}

//createImport创建导入包
func CreateImport(file *ast.File, packageName ...string) {
	var specs []ast.Spec
	i := 0
	for _, v := range packageName {
		var basic *ast.BasicLit
		if i == 0 {
			basic = &ast.BasicLit{
				ValuePos: token.Pos(2),
				Kind:     token.STRING,
				Value:    strconv.Quote(v),
			}
		} else {
			basic = &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(v),
			}
		}
		specs = append(specs, &ast.ImportSpec{
			Path: basic,
		})
		i++
	}
	if len(file.Decls) == 0 {
		ip := &ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: specs,
		}
		ips := []ast.Decl{ip}
		file.Decls = append(ips, file.Decls...)
	} else {
		var find bool
		for _, v := range file.Decls {
			if n1, ok1 := v.(*ast.GenDecl); ok1 {
				if n1.Tok == token.IMPORT {
					n1.Specs = append(n1.Specs, specs...)
					find = true
				}
			}
		}
		if !find {
			ip := &ast.GenDecl{
				Tok:   token.IMPORT,
				Specs: specs,
			}
			ips := []ast.Decl{ip}
			file.Decls = append(ips, file.Decls...)
		}
	}

}

//addImport 添加导入包
func (c *DCodeController) addImport(decl *ast.GenDecl) {
	var specs []ast.Spec
	for _, v := range needImport {
		if v.need {
			specs = append(specs, &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote(v.name),
				},
			})
		}
	}
	if len(specs) > 0 {
		decl.Specs = append(decl.Specs, specs...)
	}
}

//自动生成注释
func (c *DCodeController) genDoc(fn *ast.FuncDecl) {
	c.genTitle(fn)
	c.genDescription(fn)
	c.genRouter(fn)
	c.genParam(fn)
	c.genSuccess(fn)
	c.genFailure(fn)
}

//genSuccess 生成成功返回
func (c *DCodeController) genSuccess(fn *ast.FuncDecl) {
	if fn.Recv != nil {
		c.addComment(fn, "@Success", fmt.Sprintf("// @Success  200 {object} "))
	}
}

//genFailure 生成失败返回
func (c *DCodeController) genFailure(fn *ast.FuncDecl) {
	if fn.Recv != nil {
		c.addComment(fn, "@Failure", fmt.Sprintf("// @Failure 400 Invalid"))
	}
}

//生成标题
func (c *DCodeController) genTitle(fn *ast.FuncDecl) {
	c.addComment(fn, "@Title", fmt.Sprintf("// @Title %s 接口说明", fn.Name.Name))
}

type params struct {
	name     string
	typeStr  string
	method   string //请求方式 post,get
	defaultV string
}

//生成路由
func (c *DCodeController) genRouter(fn *ast.FuncDecl) {
	//找控制器
	if fn.Recv != nil {
		var ctl string
		if len(fn.Recv.List) > 0 {
			if n, ok := fn.Recv.List[0].Type.(*ast.StarExpr); ok {
				if n2, ok2 := n.X.(*ast.Ident); ok2 {
					c.ControllerName = n2.Name
					ctl = strings.Replace(strings.ToLower(n2.Name), "controller", "", -1)
				}
			}
		}
		if ctl != "" {
			f := FuncList{
				MethodName: strings.ToLower(fn.Name.Name),
				Method:     "any",
				Controller: ctl,
			}
			c.addComment(fn, "@router", fmt.Sprintf("// @router /%s/%s   [%s]", f.Controller, f.MethodName, f.Method))
			cm := c.getComment(fn, "@router")
			arr := strings.Fields(cm)
			if len(cm) > 1 {
				arr2 := strings.Split(arr[2], "/")
				f.MethodName = arr2[2]
				f.Controller = arr2[1]
				i1 := strings.Index(cm, "[")
				i2 := strings.Index(cm, "]")
				f.Method = lib.SubString(cm, i1+1, i2)
			}
			c.funcList = append(c.funcList, f)
		}
	}
}

//生成描述
func (c *DCodeController) genDescription(fn *ast.FuncDecl) {
	c.addComment(fn, "@Description", fmt.Sprintf("// @Description: 描述信息"))
}

var methodMap map[string]string = map[string]string{
	"Param": "body",
	"Get":   "query",
	"Post":  "formData",
}

//生成参数
func (c *DCodeController) genParam(fn *ast.FuncDecl) {
	if fn.Recv != nil {
		var p []params
		ast.Inspect(fn, func(node ast.Node) bool {
			switch node.(type) {
			case *ast.BlockStmt:
				nd := node.(*ast.BlockStmt)
				ps := c.genParamByBlockStmt(nd)
				p = append(p, ps...)
				break
			}
			return true
		})

		if len(p) > 0 {
			for _, v := range p {
				c.addComment(fn, "@Param   "+v.name, fmt.Sprintf("// @Param   %s   %s   %s   %s   \"参数说明\"", v.name, methodMap[v.method], v.typeStr, v.defaultV))
			}
		}
	}
}

//genParamByBlockStmt
func (c *DCodeController) genParamByBlockStmt(nd *ast.BlockStmt) []params {
	var p []params
	for _, v := range nd.List {
		if ck, ok := v.(*ast.AssignStmt); ok {
			//ast.Print(c.fset, ck.Rhs)
			var find bool
			var typeStr string
			var method string
			var name string
			var defaultV string
			for _, v2 := range ck.Rhs {
				ast.Inspect(v2, func(node ast.Node) bool {
					if ck2, ok2 := node.(*ast.SelectorExpr); ok2 {
						if ck2.Sel.Name == "Param" || ck2.Sel.Name == "Get" || ck2.Sel.Name == "Post" {
							if ck3, ok3 := ck2.X.(*ast.SelectorExpr); ok3 {
								if ck3.Sel.Name == "Input" {
									find = true
								}
							}
							method = ck2.Sel.Name
							//ast.Print(c.fset,ck2)
						}
					}
					return true
				})
				if n, nk := v2.(*ast.CallExpr); nk {
					if n2, nk2 := n.Fun.(*ast.SelectorExpr); nk2 {
						if n3, nk3 := n2.X.(*ast.CallExpr); nk3 {
							if len(n3.Args) > 0 {
								if n4, nk4 := n3.Args[0].(*ast.BasicLit); nk4 {
									name = strings.ReplaceAll(n4.Value, "\"", "")
								}
								if len(n3.Args) > 1 {
									if n5, nk5 := n3.Args[1].(*ast.BasicLit); nk5 {
										defaultV = strings.ReplaceAll(n5.Value, "\"", "")
									}
								}
							}
							//ast.Print(c.fset, n3.Args[0])
						}
						//ast.Print(c.fset, n2)
						typeStr = n2.Sel.Name
					}
				}
			}
			if find {
				p = append(p, params{
					name:     name,
					typeStr:  typeStr,
					method:   method,
					defaultV: defaultV,
				})
			}
		}
	}
	return p
}

//addComment 添加备注
func (c *DCodeController) addComment(fn *ast.FuncDecl, name, comment string) {
	var find bool
	if fn.Doc != nil {
		for _, v := range fn.Doc.List {
			if strings.Index(v.Text, name) != -1 {
				find = true
			}
		}
	}
	if !find {
		if fn.Doc != nil {
			fn.Doc.List = append(fn.Doc.List, &ast.Comment{
				Text:  comment,
				Slash: fn.Pos() - 1,
			})
		} else {
			cs := &ast.Comment{
				Text:  comment,
				Slash: fn.Pos() - 1,
			}
			cg := &ast.CommentGroup{
				List: []*ast.Comment{cs},
			}
			fn.Doc = cg
		}
	}
}

//getComment 获得备注信息
func (c *DCodeController) getComment(fn *ast.FuncDecl, name string) string {
	if fn.Doc != nil {
		for _, v := range fn.Doc.List {
			if strings.Index(v.Text, name) != -1 {
				return v.Text
			}
		}
	}
	return ""
}
func init() {
	needImport = getNeedImport()
}
