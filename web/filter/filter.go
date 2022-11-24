package filter
import (
	"regexp"
	"github.com/wjpxxx/letgo/web/context"
	"github.com/wjpxxx/letgo/log"
	"strings"
	"fmt"
)
//定义常量
const (
	BEFORE_STATIC =iota
	BEFORE_ROUTER
	BEFORE_EXEC
	AFTER_EXEC
	AFTER_ROUTER
)

var filterPos =map[int]string{
	BEFORE_STATIC:"before_static",
	BEFORE_ROUTER:"before_router",
	BEFORE_EXEC:"before_exec",
	AFTER_EXEC:"after_exec",
	AFTER_ROUTER:"after_router",
}

type filterData struct{
	regex *regexp.Regexp
	handler context.HandlerFunc
}
//filterMap
var filterMap map[int][]filterData

//init
func init(){
	filterMap=make(map[int][]filterData)
}

//AddFilter 添加过滤
func AddFilter(pattern string, pos int, filterFunc context.HandlerFunc){
	if _,ok:=filterPos[pos];!ok{
		log.DebugPrint("filter pos error not in BEFORE_STATIC、BEFORE_ROUTER、BEFORE_EXEC、AFTER_EXEC、AFTER_ROUTER")
		panic("filter pos error not in BEFORE_STATIC、BEFORE_ROUTER、BEFORE_EXEC、AFTER_EXEC、AFTER_ROUTER")
	}
	lasti:=strings.LastIndex(pattern,"*")
	if lasti!=-1{
		pattern=fmt.Sprintf("%s%s",pattern[:lasti],".*")
	}
	regex,regexErr:=regexp.Compile(pattern)
	if regexErr!=nil{
		log.DebugPrint("filter pattern error %v", regexErr)
		panic(regexErr)
	}
	filterMap[pos]=append(filterMap[pos], filterData{
		regex: regex,
		handler: filterFunc,
	})
}

//ExecFilter执行过滤
func ExecFilter(pos int,ctx *context.Context){
	if _,ok:=filterMap[pos];ok{
		requestPath:=strings.ToLower(ctx.Request.URL.Path)
		for _,f:=range filterMap[pos]{
			if !f.regex.MatchString(requestPath){
				continue
			}
			matches:=f.regex.FindStringSubmatch(requestPath)
			//fmt.Println(requestPath, matches,router.regex)
			if len(matches[0])!=len(requestPath) {
				continue
			}
			f.handler(ctx)
		}
	}
}