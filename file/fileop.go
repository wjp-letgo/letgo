package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/wjp-letgo/letgo/lib"
)

//BaseName 函数返回路径中的文件名部分。
//path	必需。规定要检查的路径。
//suffix	可选。规定文件扩展名。如果文件有 suffix，则不会输出这个扩展名。
func BaseName(fullPath ...string) string {
	l := len(fullPath)
	if l == 1 {
		return filepath.Base(fullPath[0])
	} else if l == 2 {
		fn := filepath.Base(fullPath[0])
		return strings.TrimSuffix(fn, fullPath[1])
	}
	return ""

}

//DirName 函数返回路径部分
//fullPath	必需。规定要检查的路径。
func DirName(fullPath string) string {
	return filepath.Dir(fullPath)
}

//BaseNameNoSuffix 函数返回路径中的文件名部分,不带扩展名。
//fullPath	必需。规定要检查的路径。
func BaseNameNoSuffix(fullPath string) string {
	fn := filepath.Base(fullPath)
	suffix := filepath.Ext(fullPath)
	return strings.TrimSuffix(fn, suffix)
}

//GetExt 获得文件后缀，扩展名
func GetExt(fullPath string) string {
	suffix := filepath.Ext(fullPath)
	return suffix
}

//GetFilesName 获取目录下的所有文件名列表
func GetFilesName(pathName string) []string {
	dir, err := ioutil.ReadDir(pathName)
	if err != nil {
		return nil
	}
	var list []string
	for _, path := range dir {
		if !path.IsDir() {
			list = append(list, path.Name())
		}
	}
	return list
}

//Mkdir 创建目录
func Mkdir(path string) {
	_, e1 := os.Stat(path)
	if e1 != nil && !os.IsExist(e1) {
		if path != "" {
			os.MkdirAll(path, 0777)
		}
	}
}

//FileExist 判断文件是否存在，存在返回true 不存在返回false
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true
	} else {
		return false
	}
}
//Slash 根据系统返回正常的路径
func Slash(path string) string{
	if runtime.GOOS=="windows" {
		return strings.ReplaceAll(path, "/", string(os.PathSeparator))
	}else{
		return strings.ReplaceAll(path, "\\", string(os.PathSeparator))
	}
}
//SlashR  将'\'转为'/'
func SlashR(path string)string{
	return strings.ReplaceAll(path, "\\", "/")
}

//Filer 文件接口
type Filer interface{
	Size() int64
	ReadAll() ([]byte, error)
	ReadLine() ([]byte, error)
	ReadBlock(size int64) ([]byte, int64)
	Write(b []byte)
	WriteAt(b []byte, offset int64) int
	WriteAppend(b []byte)
	Truncate()
	Ext() string
	Name() string
	Path() string
	FullPath()string
	ModifyTime() int
}
//File 文件类
type File struct {
	name string
	path string
	fullPath string
	handle *os.File
	seek int64
}
//setFullPath 设置文件路径
func (f *File)setFullPath(fullPath string) *File{
	f.fullPath=fullPath
	return f
}
//Open 打开文件
func (f *File) open() {
	f.name=BaseName(f.fullPath)
	f.path=DirName(f.fullPath)
	_, err := os.Stat(f.path)
	if err != nil && !os.IsExist(err) {
		if f.path != "" {
			Mkdir(f.path)
		}
	}
	_, err = os.Stat(f.fullPath)
	if err == nil || os.IsExist(err) {
		f.handle, err = os.OpenFile(f.fullPath, os.O_RDWR, 0644)
	} else {
		f.handle, err = os.Create(f.fullPath)
	}
	if err!=nil{
		return
	}
}
//Close 关闭文件
func (f *File) close() {
	f.handle.Close()
}
//Size 获得文件大小
func (f *File) Size() int64 {
	f.open()
	defer f.close()
	stat, er := f.handle.Stat()
	if er == nil {
		return stat.Size()
	}
	return -1
}
//ReadAll 读取全部
func (f *File) ReadAll() ([]byte, error) {
	f.open()
	defer f.close()
	return ioutil.ReadAll(f.handle)
}
//ReadLine 读取一行
func (f *File) ReadLine() ([]byte, error) {
	f.open()
	defer f.close()
	reader:=bufio.NewReader(f.handle)
	line, _, err := reader.ReadLine()
	return line,err
}
//ReadBlock 读取块
func (f *File) ReadBlock(size int64) ([]byte, int64) {
	seek:=f.seek+size
	var buf []byte
	if f.Size()-seek>=size{
		buf = make([]byte, size)
	} else if f.Size()-seek > 0 {
		buf = make([]byte, f.Size()-seek)
	}else if f.Size()-seek <= 0 &&f.Size()-f.seek>=0 {
		buf = make([]byte, f.Size()-f.seek)
		seek=f.Size()
	}else {
		return nil, -1
	}
	f.open()
	defer f.close()
	_, err := f.handle.ReadAt(buf, f.seek)
	if err != nil {
		return nil, -2
	}
	f.seek=seek
	return buf, seek
}
//Write 写入
func (f *File) Write(b []byte) {
	f.open()
	defer f.close()
	f.handle.Write(b)
	l:=int64(len(b))
	f.handle.Truncate(l)
}
//WriteAt 在固定点写入
func (f *File) WriteAt(b []byte, offset int64) int{
	f.open()
	defer f.close()
	seek,_:=f.handle.WriteAt(b,offset)
	l:=int64(len(b))
	f.handle.Truncate(l+offset)
	return seek
}
//WriteAppend 追加文件
func (f *File) WriteAppend(b []byte) {
	f.open()
	defer f.close()
	size:=f.Size()
	if size>-1{
		f.WriteAt(b,size)
	}
}
//Truncate 清空文件
func (f *File) Truncate() {
	f.open()
	defer f.close()
	f.handle.Truncate(0)
}
//Ext 文件后缀名
func (f *File) Ext() string {
	return GetExt(f.fullPath)
}
//Name 文件名
func (f *File) Name() string {
	return f.name
}
//Path 文件路径
func (f *File) Path() string {
	return f.path
}
//FullPath 文件完整路径
func (f *File) FullPath() string {
	return f.fullPath
}
//ModifyTime 文件修改时间
func (f *File) ModifyTime() int {
	f.open()
	defer f.close()
	stat, er := f.handle.Stat()
	if er == nil {
		return lib.TimeByTime(stat.ModTime()) 
	}
	return -1
}
//NewFile 新建文件
func NewFile(fullPath string) Filer{
	var file File
	return file.setFullPath(fullPath)
}
//PutContent 往文件写入数据
func PutContent(fullPath,content string) {
	f:=NewFile(fullPath)
	f.Write([]byte(content))
}

//PutContentBytes 往文件写入数据
func PutContentBytes(fullPath string,content []byte) {
	f:=NewFile(fullPath)
	f.Write(content)
}

//PutContentAppend 往文件写入数据以追加的 形式
func PutContentAppend(fullPath,content string) {
	f:=NewFile(fullPath)
	f.WriteAppend([]byte(content))
}
//PutContentBytesAppend 往文件写入数据以追加的 形式
func PutContentBytesAppend(fullPath string,content []byte) {
	f:=NewFile(fullPath)
	f.WriteAppend(content)
}
//GetContent 获得文件内容
func GetContent(fullPath string) string{
	f:=NewFile(fullPath)
	content,_:=f.ReadAll()
	return string(content)
}