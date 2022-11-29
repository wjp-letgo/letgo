package limiting

import (
	"sort"

	"github.com/wjp-letgo/letgo/lib"
)

//计数器算法
//FlowCounter
type FlowCounter struct{
	counter int //当前请求数量
	counterLimit int //时间窗口内最大请求数
	firstTime int64 //起始时间
	window int64 	//时间窗口ms
}

//Pass 是否达到限流 true 达到限流条件 false 还未达到限流
func (f *FlowCounter)Pass() bool {
	now:=lib.TimeLong()
	if now<f.firstTime+f.window {
		//在窗范围内
		f.counter++
		return f.counter>f.counterLimit
	}else{
		f.firstTime=now
		f.counter=1
		return false
	}
}
//NewFlowCounter
func NewFlowCounter(counterLimit int, window int64)*FlowCounter{
	return &FlowCounter{0,counterLimit,0,window}
}
//NewFlowCounterByParams
func NewFlowCounterByParams(args ...interface{})*FlowCounter{
	if len(args)==1{
		return NewFlowCounter(lib.InterfaceToInt(args[0]), 1000)
	}else if len(args)==2{
		return NewFlowCounter(lib.InterfaceToInt(args[0]),lib.InterfaceToInt64(args[1]))
	}else{
		return NewFlowCounter(200, 1000)
	}
}

//计数器算法-滑动窗口
//FlowRollingCounter
type FlowRollingCounter struct{
	counterLimit int //时间窗口内最大请求数
	subWinNums int   //子窗口数量
	firstTime int64 //起始时间
	window int64   //时间窗口ms
	dis float32
	subWins subWins //子窗口
}
//subWins
type subWins []subWin
//subWin 子窗口
type subWin struct{
	start int64 //起始时间
	end int64   //结束时间
	counter int //当前请求数量
}
//Len
func (s subWins) Len() int {
	return len(s)
}
//Swap
func (s subWins) Swap(i, j int){
	s[i],s[j]=s[j],s[i]
}

//Less
func (s subWins) Less(i, j int) bool {
	return s[i].start<s[j].start
}

//Init
func (f *FlowRollingCounter)Init()*FlowRollingCounter{
	f.firstTime=lib.TimeLong()
	f.dis=float32(f.window)/float32(f.subWinNums)
	for i:=0;i<f.subWinNums;i++{
		start:=f.firstTime+int64(i)*int64(f.dis)
		sub:=subWin{
			start:start,
			end: start+int64(f.dis)-1,
			counter:0,
		}
		f.subWins=append(f.subWins,sub)
	}
	sort.Sort(f.subWins)
	return f
}

//Pass 是否达到限流 true 达到限流条件 false 还未达到限流
func (f *FlowRollingCounter)Pass() bool {
	now:=lib.TimeLong()
	f.moveLast(now)
	f.addCounter(now)
	//fmt.Println(lib.TimeLongToStr(now),f.subWins)
	return f.sumCounter()>f.counterLimit
}
//addCounter
func (f *FlowRollingCounter)addCounter(now int64) {
	for i,sub:=range f.subWins{
		if now>=sub.start&&now<=sub.end{
			f.subWins[i].counter++
		}
	}
}

//moveLast
func (f *FlowRollingCounter)moveLast(now int64){
	for i,sub:=range f.subWins{
		if now>sub.end{
			f.subWins=f.subWins[i+1:]
			start:=f.subWins[len(f.subWins)-1].end+int64(1)
			f.subWins=append(f.subWins, subWin{
				start:start,
				end:start+int64(f.dis)-1,
				counter: 0,
			})
			f.moveLast(now)
			break
		}
	}
}

//sumCounter
func (f *FlowRollingCounter)sumCounter()int{
	var sum int=0
	for _,sub:=range f.subWins{
		sum+=sub.counter
	}
	return sum
}

//NewFlowCounter
//counterLimit:总限制数量
//subWinNums:子窗口个数
//window:窗口大小 这个时间范围内请求个数不得超过counterLimit
func NewFlowRollingCounter(counterLimit,subWinNums int, window int64)*FlowRollingCounter{
	f:= &FlowRollingCounter{counterLimit,subWinNums,0,window,0,nil}
	return f.Init()
}
//NewFlowRollingCounterByParams
func NewFlowRollingCounterByParams(args ...interface{})*FlowRollingCounter{
	if len(args)==1{
		return NewFlowRollingCounter(lib.InterfaceToInt(args[0]), 10, 1000)
	}else if len(args)==2{
		return NewFlowRollingCounter(lib.InterfaceToInt(args[0]),lib.InterfaceToInt(args[1]),1000)
	}else if len(args)==3{
		return NewFlowRollingCounter(lib.InterfaceToInt(args[0]),lib.InterfaceToInt(args[1]),lib.InterfaceToInt64(args[2]))
	}else{
		return NewFlowRollingCounter(200, 10, 1000)
	}
}