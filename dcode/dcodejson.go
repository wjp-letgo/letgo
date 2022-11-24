package dcode

import (
	"fmt"
	"github.com/wjpxxx/letgo/lib"
	"go/ast"
	"go/token"
	"strings"
)

//路由检测
type DCodeJson struct {
	fset    *token.FileSet
	mfile   *ast.File
	hasJson bool
	structs []entity
}

type entity struct {
	name string
	end  token.Pos
}

//Visit 遍历抽象语法树
func (c *DCodeJson) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.GenDecl:
		genDecl := node.(*ast.GenDecl)
		c.genDecl(genDecl)
		break
	}
	return c
}

//Finish 遍历完成
func (c *DCodeJson) Finish() {
	for _, fname := range c.structs {
		var find bool
		ast.Inspect(c.mfile, func(node ast.Node) bool {
			switch node.(type) {
			case *ast.FuncDecl:
				fn := node.(*ast.FuncDecl)
				if fn.Name.Name == "String" && len(fn.Recv.List) > 0 {
					if tp, ok := fn.Recv.List[0].Type.(*ast.StarExpr); ok {
						if x, ok2 := tp.X.(*ast.Ident); ok2 {
							if x.Name == fname.name {
								find = true
							}
						}
					}
				}
				break
			}
			return true
		})
		if !find {
			c.createJsonStringFunc(fname)
		}
	}
	if c.hasJson {
		CreateImport(c.mfile, "github.com/wjpxxx/letgo/lib")
	}
}

//createJsonStringFunc 创建json string函数
func (c *DCodeJson) createJsonStringFunc(fname entity) {
	var fieldList []*ast.Field
	var names []*ast.Ident
	names = append(names, &ast.Ident{
		Name: strings.ToLower(fname.name),
		Obj: &ast.Object{
			Kind: ast.Var,
			Name: strings.ToLower(fname.name),
			Decl: &ast.Field{},
		},
	})
	fieldList = append(fieldList, &ast.Field{
		Names: names,
		Type: &ast.StarExpr{
			X: &ast.Ident{
				Name: fname.name,
				Obj: &ast.Object{
					Kind: ast.Typ,
					Name: fname.name,
					Decl: &ast.TypeSpec{},
				},
			},
		},
	})
	var lst2 []*ast.Field
	lst2 = append(lst2, &ast.Field{
		Type: &ast.Ident{
			Name: "string",
		},
	})
	var lst3 []ast.Stmt
	var lst4 []ast.Expr
	var lst5 []ast.Expr
	lst5 = append(lst5, &ast.Ident{
		Name: strings.ToLower(fname.name),
	})
	lst4 = append(lst4, &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X: &ast.Ident{
				Name: "lib",
			},
			Sel: &ast.Ident{
				Name: "ObjectToString",
			},
		},
		Args: lst5,
	})
	lst3 = append(lst3, &ast.ReturnStmt{
		Results: lst4,
	})
	c.mfile.Decls = append(c.mfile.Decls, &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: fieldList,
		},
		Name: &ast.Ident{
			Name: "String",
		},
		Type: &ast.FuncType{
			Results: &ast.FieldList{
				List: lst2,
			},
		},
		Body: &ast.BlockStmt{
			List: lst3,
		},
	})
}

//genDecl 字节点
func (c *DCodeJson) genDecl(decl *ast.GenDecl) {
	if decl.Doc != nil {
		var isjson int
		for _, v := range decl.Doc.List {
			if v.Text == "//@json" {
				isjson = 1
			} else if v.Text == "//@Json" {
				isjson = 2
			}
		}
		if isjson > 0 {
			c.hasJson = true
			for _, v1 := range decl.Specs {
				if vv, ok := v1.(*ast.TypeSpec); ok {
					if t, ok2 := vv.Type.(*ast.StructType); ok2 {
						c.structs = append(c.structs, entity{
							name: vv.Name.Name,
							end:  vv.End(),
						})
						for _, f := range t.Fields.List {
							var name string
							if len(f.Names) > 0 {
								name = f.Names[0].Name
							}
							var js string
							if isjson == 1 {
								js = fmt.Sprintf("`json:\"%s\"`", lib.UnderLineName(name))
							} else if isjson == 2 {
								js = fmt.Sprintf("`json:\"%s\"`",lib.FirstToLower(name))
							}
							if f.Tag == nil {
								f.Tag = &ast.BasicLit{
									Kind:  token.STRING,
									Value: js,
								}
							} else {
								f.Tag = &ast.BasicLit{
									Kind:  token.STRING,
									Value: js,
								}
							}
						}
					}
				}
			}
		}
	}
}
