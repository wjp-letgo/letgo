package command

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wjp-letgo/letgo/lib"
)


func TestCommand(t *testing.T) {
	cmd:=New().Cd("D:\\Development\\go\\web\\src").SetCMD("dir")
	//cmd.AddPipe(New().SetCMD("find","'\\c'","'80'"))
	cmd.Run()
	cmd2:=New().Cd("D:\\").SetCMD("dir")
	var b bytes.Buffer
	cmd2.SetStdout(&b)
	cmd2.Run()
	fmt.Println("==========================")
	fmt.Println(lib.Gb2312ToUtf8(string(b.Bytes())))
}