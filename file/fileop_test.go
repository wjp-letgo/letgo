package file

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T){
	name:="upload/2021/a.config"
	PutContentAppend(name, "sssssssss");
	fmt.Println(GetContent(name))
}