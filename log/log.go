package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/wjp-letgo/letgo/file"
	"github.com/wjp-letgo/letgo/lib"
)


var config LogConfig

//IsDebug
func IsDebug()bool{
	return config.Debug
}
//GetWriter
func GetWriter()io.Writer{
	var w io.Writer
	switch config.Writer {
	case "os":
		w=os.Stdout
	case "file":
		path:=file.DirName(config.LogFilePath)
		if path=="."{
			path=config.LogFilePath
		}
		file.Mkdir(path)
		config.File,_=os.OpenFile(path+"/debug.log", os.O_WRONLY|os.O_CREATE, 0666)
		w=config.File
	default:
		panic("unknown writer, The only options are file or os")
	}
	return w
}
//Close
func Close(){
	if config.Writer=="file" {
		config.File.Close()
	}
}

//DebugPrint
func DebugPrint(format string,values ...interface{}){
	if IsDebug() {
		w:=GetWriter()
		defer Close()
		if !strings.HasSuffix(format,"\n") {
			format+="\n"
		}
		timeStr:=lib.Now()
		fmt.Fprintf(w,"[Letgo-debug:"+timeStr+"]"+format,values...)
	}
}

//PanicPrint
func PanicPrint(format string,values ...interface{}){
	if IsDebug() {
		if !strings.HasSuffix(format,"\n") {
			format+="\n"
		}
		timeStr:=lib.Now()
		panic(fmt.Sprintf("[Letgo-debug:"+timeStr+"]"+format,values...))
	}
}

//init
func init() {
	logFile:="config/log.config"
	cfgFile:=file.GetContent(logFile)
	if cfgFile==""{
		logConfig:=LogConfig{
			Debug: true,
			Writer: "os",
			LogFilePath:"",//when Writer is file
		}
		file.PutContent(logFile,fmt.Sprintf("%v",logConfig))
		panic("please setting log config in config/log.config file")
	}
	lib.StringToObject(cfgFile,&config)
}