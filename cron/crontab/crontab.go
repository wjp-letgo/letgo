package crontab

import (
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/wjp-letgo/letgo/cron/context"
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/log"
)

//  * * * * * *   秒钟(0-59) 分钟(0-59) 小时(1-23) 日期(1-31) 月份(1-12) 星期(0-6)
var (
	//globalTaskManager
	globalTaskManager *taskManager
	seconds =trange{0,59,nil}
	minutes =trange{0,59,nil}
	hours =trange{0,23,nil}
	days =trange{1,31,nil}
	months =trange{1,12,map[string]uint{
		"jan": 1,
		"feb": 2,
		"mar": 3,
		"apr": 4,
		"may": 5,
		"jun": 6,
		"jul": 7,
		"aug": 8,
		"sep": 9,
		"oct": 10,
		"nov": 11,
		"dec": 12,
	}}
	weeks =trange{0,6,map[string]uint{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}}
)
//trange 取值范围
type trange struct{
	start,end uint
	title map[string]uint
}
//taskManager
type taskManager struct{
	taskLock sync.RWMutex
	taskList map[string]*task
	started       bool
	stop chan bool
	wait sync.WaitGroup
}
//DoFunc
type DoFunc func(*context.Context)

//Task
type task struct{
	context *context.Context
	specStr string
	spec	CronInfo
	do DoFunc
	next int
}
//CronInfo
type CronInfo struct{
	Second uint64
	Minute uint64
	Hour   uint64
	Day    uint64
	Month  uint64
	Week   uint64
}

//newTaskManager
func newTaskManager()*taskManager{
	return &taskManager{
		taskList:make(map[string]*task),
		started: false,
		stop:make(chan bool),
	}
}

//init
func init(){
	globalTaskManager=newTaskManager()
}

//taskMapSort
type taskMapSort struct{
	Keys []string
	Values []*task
}
//newMapSort
func newMapKV(taskList map[string]*task)taskMapSort{
	ms:=taskMapSort{
		Keys:make([]string, 0,len(taskList)),
		Values: make([]*task, 0,len(taskList)),
	}
	for k,_:=range taskList{
		ms.Keys=append(ms.Keys, k)
	}
	sort.Strings(ms.Keys)
	for _,v:=range ms.Keys{
		ms.Values=append(ms.Values, taskList[v])
	}
	return ms
}

//managerWait
func (t *taskManager)managerWait(){
	for {
		if t.started==false{
			break
		}
		time.Sleep(100* time.Millisecond)
	}
	t.wait.Wait()
}
//managerStop
func (t *taskManager)managerStop(){
	t.stop<-true
	t.taskLock.Lock()
	if t.started {
		t.started=false
	}
	t.taskLock.Unlock()
}

//Start
func (t *taskManager)Start(){
	t.taskLock.Lock()
	defer t.taskLock.Unlock()
	if t.started{
		return
	}
	t.started=true
	go t.run()
}

//run
func (t *taskManager)run(){
	sortList:=newMapKV(t.taskList)
	for{
		select{
			default:
				t.startTask(sortList)
			case <-t.stop:
				return
		}
	}
}

//startTask
func (t *taskManager)startTask(sortList taskMapSort){
	for _,tsk:=range sortList.Values{
		if timeToStart(tsk)&&!tsk.context.IsDone(){
			tsk.context.SetDone(true)
			t.wait.Add(1)
			go func(tk *task){
				defer t.wait.Done()
				tsk.context.Now=lib.Time()
				tk.do(tk.context)
				tk.context.SetDone(false)
			}(tsk)
		}
	}
}
//timeToStart 时间是否到了可以启动
func timeToStart(tsk *task)bool{
	now :=time.Now().Local()
	tnow:=lib.TimeByTime(now)
	if tsk.next==0{
		tsk.next=next(now, tsk)
	} else if tnow>=tsk.next{
		//log.DebugPrint("now(%s)=next(%s) %t",lib.TimeToStr(lib.TimeByTime(now)),lib.TimeToStr(tsk.next),tnow>=tsk.next)
		tsk.next=next(now, tsk)
		return true
	}
	//log.DebugPrint("now(%s)=next(%s)",lib.TimeToStr(lib.TimeByTime(now)),lib.TimeToStr(tsk.next))
	return false
}
//next 下次运行时间
func next(t time.Time, tsk *task) int{
	t = t.Add(1*time.Second - time.Duration(t.Nanosecond())*time.Nanosecond)
	added := false
	yearLimit := t.Year() + 5
WRAP:
	if t.Year() > yearLimit {
		return lib.Time()
	}
	for 1<<uint(t.Month())&tsk.spec.Month == 0 {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
		}
		t = t.AddDate(0, 1, 0)
		if t.Month() == time.January {
			goto WRAP
		}
	}
	for !dayMatches(tsk, t) {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		}
		t = t.AddDate(0, 0, 1)

		if t.Day() == 1 {
			goto WRAP
		}
	}
	for 1<<uint(t.Hour())&tsk.spec.Hour == 0 {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
		}
		t = t.Add(1 * time.Hour)

		if t.Hour() == 0 {
			goto WRAP
		}
	}
	for 1<<uint(t.Minute())&tsk.spec.Minute == 0 {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
		}
		t = t.Add(1 * time.Minute)

		if t.Minute() == 0 {
			goto WRAP
		}
	}
	for 1<<uint(t.Second())&tsk.spec.Second == 0 {
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
		}
		t = t.Add(1 * time.Second)

		if t.Second() == 0 {
			goto WRAP
		}
	}
	return lib.TimeByTime(t)
}

func dayMatches(tsk *task, t time.Time) bool {
	var (
		domMatch = 1<<uint(t.Day())&tsk.spec.Day > 0
		dowMatch = 1<<uint(t.Weekday())&tsk.spec.Week > 0
	)
	var starBit uint64 = 1 << 63
	if tsk.spec.Day&starBit > 0 || tsk.spec.Week&starBit > 0 {
		return domMatch && dowMatch
	}
	return domMatch || dowMatch
}

//addCron
func (t *taskManager)addCron(name string,spec string, f DoFunc){
	t.taskLock.Lock()
	defer t.taskLock.Unlock()
	ctx:=context.NewContext()
	ctx.Name=name
	ctx.TaskNo=0
	ctx.Done=false
	t.taskList[name]=&task{
		context:ctx,
		specStr:spec,
		spec: t.getCronInfo(spec),
		do:f,
	}
}
//getCronInfo
func (t *taskManager)getCronInfo(spec string)CronInfo{
	fields:=strings.Fields(spec)
	if len(fields)!=6{
		log.PanicPrint("Expected 6 fields, found %d: %s", len(fields), spec)
	}
	return CronInfo{
		Second:getValue(fields[0], seconds),
		Minute: getValue(fields[1], minutes),
		Hour: getValue(fields[2], hours),
		Day: getValue(fields[3], days),
		Month: getValue(fields[4],months),
		Week: getValue(fields[5], weeks),
	}
}
//Start start task
func Start(){
	globalTaskManager.Start()
}

//Wait wait all task done
func Wait(){
	globalTaskManager.managerWait()
}

//Stop stop task
func Stop(){
	globalTaskManager.managerStop()
}
//StartAndWait  start and wait all task done
func StartAndWait(){
	Start()
	Wait()
}

//AddCron 添加定时任务
func AddCron(name string,spec string, f DoFunc){
	globalTaskManager.addCron(name,spec,f)
}
//getValue
func getValue(name string, r trange) uint64{
	var bits uint64=1<<63
	rs:=strings.FieldsFunc(name, func(rr rune) bool{return rr==','})
	for _,v:=range rs{
		bits|=getRange(v,r)
	}
	return bits
}
//getRange
func getRange(v string, r trange)uint64{
	var start,end,step uint
	stepArr:=strings.Split(v, "/")
	startAndEnd:=strings.Split(stepArr[0], "-")
	if startAndEnd[0]=="*"||startAndEnd[0]=="?"{
		start=r.start
		end=r.end
	} else {
		start=parseIntOrTitle(startAndEnd[0], r.title)
		switch len(startAndEnd) {
		case 1:
			end=start
		case 2:
			end=parseIntOrTitle(startAndEnd[1],r.title)
		default:
			log.PanicPrint("Too many -:%s",v)
		}
	}
	switch len(stepArr) {
	case 1:
		step=1
	case 2:
		step=lib.StrToUInt(stepArr[1])
	default:
		log.PanicPrint("Too many /:%s",v)
	}
	if start<r.start{
		log.PanicPrint("start (%d) below (%d): %s",start, r.start, v)
	}
	if end>r.end{
		log.PanicPrint("end :(%d) above (%d): %s",end, r.end, v)
	}
	if start > end {
		log.PanicPrint("start (%d) beyond end of range (%d): %s", start, end, v)
	}
	bits:=getBit(start, end, step)
	return bits
}
//parseIntOrTitle
func parseIntOrTitle(name string, title map[string]uint)uint{
	if title!=nil{
		if v,ok:=title[strings.ToLower(name)];ok{
			return v
		}
	}
	return lib.StrToUInt(name)
}

//getBit 获得位比特， int64 的二进制来标记分钟上的集合,从右往左第1位标识0分钟，第2位标记标识1分钟一直到第60位标记第59分钟
func getBit(start,end,step uint)uint64{
	var bits uint64
	for i:=start;i<=end;i+=step{
		bits|=1<<i
	}
	return bits
}