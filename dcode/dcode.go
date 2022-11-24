package dcode

import (
	"bytes"
	"github.com/wjpxxx/letgo/file"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//ControllerInfo 控制器信息
type ControllerInfo struct {
	PackageName    string
	Path           string
	Funcs          []FuncList
	ControllerName string
}

//自动检测代码控制器
func DcodeController(controllerDir string) []ControllerInfo {
	var cs []ControllerInfo
	filepath.Walk(controllerDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if file.GetExt(path) == ".go" {
				fset := token.NewFileSet()
				fl, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
				if err == nil {
					//ast.Print(fset,file)
					dc := &DCodeController{fset: fset, mfile: fl}
					ast.Walk(dc, fl)
					dc.Finish()
					var buf []byte
					bf := bytes.NewBuffer(buf)
					format.Node(bf, fset, fl)
					//wbf,_:=format.Source(bf.Bytes())
					ioutil.WriteFile(path, bf.Bytes(), 0644)
					//fmt.Println(bf.String())
					cs = append(cs, ControllerInfo{
						PackageName:    dc.PackageName,
						Path:           strings.ReplaceAll(file.DirName(path), "\\", "/"),
						Funcs:          dc.GetFuncList(),
						ControllerName: dc.ControllerName,
					})
				}
			}
		}
		return nil
	})
	return cs
}

//自动检测代码路由
func DcodeRouter(router, controllerDir string) bool {
	if file.GetExt(router) == ".go" {
		md := GetGoModName()
		rs := DcodeController(controllerDir)
		fset := token.NewFileSet()
		fl, err := parser.ParseFile(fset, router, nil, parser.ParseComments)
		if err == nil {
			dc := &DCodeRouter{fset: fset, mfile: fl, modName: md, controller: rs}
			ast.Walk(dc, fl)
			dc.Finish()
			var buf []byte
			bf := bytes.NewBuffer(buf)
			format.Node(bf, fset, fl)
			//wbf,_:=format.Source(bf.Bytes())
			ioutil.WriteFile(router, bf.Bytes(), 0644)
		}
		return true
	}
	return false
}

func DcodeJson(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if file.GetExt(path) == ".go" {
				fset := token.NewFileSet()
				fl, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
				if err == nil {
					//ast.Print(fset,file)
					dc := &DCodeJson{fset: fset, mfile: fl}
					ast.Walk(dc, fl)
					dc.Finish()
					var buf []byte
					bf := bytes.NewBuffer(buf)
					format.Node(bf, fset, fl)
					//wbf,_:=format.Source(bf.Bytes())
					ioutil.WriteFile(path, bf.Bytes(), 0644)
					//fmt.Println(bf.String())
				}
			}
		}
		return nil
	})
}

func GetGoModName() string {
	f := file.NewFile("./go.mod")
	l, _ := f.ReadLine()
	if len(l) > 0 {
		arr := strings.Split(string(l), " ")
		return strings.Trim(arr[1], " ")
	}
	return ""
}
