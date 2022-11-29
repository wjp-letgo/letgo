package walkdir

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/wjp-letgo/letgo/file"
	"github.com/wjp-letgo/letgo/log"
)

//Walk 遍历目录
func Walk(dirname string,options *Options){
	if !filepath.IsAbs(dirname) {
		dirname,_=filepath.Abs(dirname)
	}
	//dirname=filepath.FromSlash(dirname)
	dir,err:=ioutil.ReadDir(dirname)
	if err!=nil{
		log.DebugPrint("walk dir error %s", err)
		return 
	}
	for _,p:=range dir{
		if p.IsDir() {
			tmpDirName:=filepath.Join(dirname,p.Name())
			Walk(tmpDirName, options)
		}else{
			fullPath:=filepath.Join(dirname,p.Name())
			if options.Callback!=nil&&filter(fullPath,options){
				options.Callback(dirname,p.Name(),fullPath,options.LocationPath,options.RemotePath)
			}
		}
	}
}
type WalkFunc func(pathName,fileName,fullName,LocationPath,RemotePath string)
//Options
type Options struct{
	Callback WalkFunc
	Filter []string //Filter files
	LocationPath string
	RemotePath string
}
//filter
func filter(fullPath string,options *Options)bool{
	filterArray:=options.Filter
	localPath:=options.LocationPath
	fullPath=file.SlashR(fullPath)
	for _,f:=range filterArray{
		if !filepath.IsAbs(f) {
			f=filepath.Join(localPath,f)
		}
		f=file.SlashR(f)
		lasti:=strings.LastIndex(f,"*")
		if lasti!=-1{
			f=fmt.Sprintf("%s%s",f[:lasti],".*")
		}
		//log.DebugPrint("f %s",f)
		//log.DebugPrint("f abs: %s %s",f,fullPath)
		regex,err:=regexp.Compile(f)
		if err!=nil{
			log.DebugPrint("filter error %v",err)
			continue
		}
		//log.DebugPrint("regex:%v",regex.MatchString(fullPath))
		if regex.MatchString(fullPath){
			return false
		}
	}
	return true
}