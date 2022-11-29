package log

import (
	"os"

	"github.com/wjp-letgo/letgo/lib"
)

type LogConfig struct{
	Debug bool `json:"debug"`
	Writer string `json:"writer"`
	LogFilePath string `json:"logFilePath"`
	File *os.File
}

func (l LogConfig)String()string{
	return lib.ObjectToString(l)
}