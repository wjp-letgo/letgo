package walkdir

import (
	"fmt"
	"testing")


func TestWalkDir(t *testing.T){
	Walk("./",&Options{
		Callback:func(path,file,fullName string){
			fmt.Println(path,file,fullName)
		},
		Filter: []string{"*"},
	})
}