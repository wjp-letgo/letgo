package task

import (
	"fmt"
	"testing"
	"time"

	"github.com/wjp-letgo/letgo/cron/context"
)
func TestTask(t *testing.T) {
	RegisterTaskByMethodAndFilter("add",3,func(ctx *context.Context){
		fmt.Println(ctx.Name,ctx.TaskNo)
		//fmt.Println()
		//time.Sleep(1*time.Second)
	},func(ctx *context.Context)bool{
		time.Sleep(1*time.Second)
		fmt.Println("判断",ctx.Name)
		return false
	})
	go func(){
		time.Sleep(105*time.Second)
		Stop()
	}()
	StartAndWait()
}