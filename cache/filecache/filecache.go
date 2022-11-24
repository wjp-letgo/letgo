package filecache

import (
	"github.com/wjpxxx/letgo/cache/icache"
	"github.com/wjpxxx/letgo/encry"
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/lib"
	"fmt"
	"os"
	"strings"
)


//FileCache 文件缓存
type FileCache struct {
	path string
}
//SetPath 设置缓存文件的路径
func (f *FileCache)SetPath(path string) icache.ICacher{
	f.path=path
	file.Mkdir(path)
	return f
}
//Set
func (f *FileCache)Set(key string, value interface{}, overtime int64) bool {
	fullPath:=f.getFullName(key)
	if overtime>-1{
		file.PutContent(fullPath, fmt.Sprintf("%d#",lib.Time()+int(overtime))+string(lib.Serialize(value)))
	} else {
		file.PutContent(fullPath, fmt.Sprintf("%d#",int(overtime))+string(lib.Serialize(value)))
	}
	return true
}
func (f *FileCache)getFullName(key string) string{
	name:=encry.Md5(key)
	fullPath:=fmt.Sprintf("%s%s",f.path,name)
	return fullPath
}
//Get
func (f *FileCache)Get(key string, value interface{}) bool {
	fullPath:=f.getFullName(key)
	if !file.FileExist(fullPath) {
		return false
	}
	content:=file.GetContent(fullPath)
	if (len(content)==0) {
		return false
	}
	i:=strings.Index(content,"#")
	overTime:=lib.SubString(content,0,i)
	overTimeInt:=(&lib.Data{Value: overTime}).Int()
	if overTimeInt>-1&&lib.Time()>overTimeInt {
		//过期了
		os.Remove(fullPath)
		return false
	}
	cbyte:=[]byte(content)
	serializeContent:=cbyte[i+1:]
	lib.UnSerialize([]byte(serializeContent),value)
	return true
}
//Del
func (f *FileCache)Del(key string) bool {
	fullPath:=f.getFullName(key)
	err:=os.Remove(fullPath)
	if err!=nil{
		return false
	}
	return true
}
//FlushDB
func (f *FileCache)FlushDB() bool {
	err:=os.RemoveAll(f.path)
	if err!=nil{
		return false
	}
	return true
}
//NewFileCache 文件缓存
func NewFileCache() icache.ICacher{
	return NewFileCacheByPath("runtime/cache/")
}
//NewFileCacheByPath
func NewFileCacheByPath(path string)icache.ICacher{
	f:=&FileCache{}
	return f.SetPath(path)
}