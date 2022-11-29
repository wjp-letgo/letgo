package task

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/wjp-letgo/letgo/cron/context"
	"github.com/wjp-letgo/letgo/lib"
)

//DoFunc
type DoFunc func(*context.Context)

//Task
type task struct{
	context *context.Context
	do DoFunc
	filter func(*context.Context)bool
}

//globalTaskManager
var globalTaskManager *taskManager

//taskManager
type taskManager struct{
	taskLock sync.RWMutex
	taskList map[string]task
	started       bool
	stop chan bool
	wait sync.WaitGroup
}

//taskMapSort
type taskMapSort struct{
	Keys []string
	Values []task
}

//newMapSort
func newMapKV(taskList map[string]task)taskMapSort{
	ms:=taskMapSort{
		Keys:make([]string, 0,len(taskList)),
		Values: make([]task, 0,len(taskList)),
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

//newTaskManager
func newTaskManager()*taskManager{
	return &taskManager{
		taskList:make(map[string]task),
		started: false,
		stop:make(chan bool),
	}
}

//init
func init(){
	globalTaskManager=newTaskManager()
}
//RegisterTask
func (t *taskManager)RegisterTask(name string,taskNums int,call DoFunc){
	t.registerTask(name,taskNums,call, nil)
}
//RegisterTask
func (t *taskManager)RegisterTaskByFilter(name string,taskNums int,call DoFunc,filter func(*context.Context)bool){
	t.registerTask(name,taskNums,call, filter)
}
//RegisterTask
func (t *taskManager)registerTask(name string,taskNums int,call DoFunc,filter func(*context.Context)bool){
	t.taskLock.Lock()
	for i:=0;i<taskNums;i++{
		ctx:=context.NewContext()
		ctx.Name=name
		ctx.TaskNo=i
		ctx.Done=false
		key:=fmt.Sprintf("%s-%d",name,i)
		t.taskList[key]=task{
			context:ctx,
			do:call,
			filter: filter,
		}
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
			time.Sleep(10*time.Millisecond)
		case <-t.stop:
			return
		}
	}
}
//startTask
func (t *taskManager)startTask(sortList taskMapSort){
	for _,tsk:=range sortList.Values{
		if !tsk.context.IsDone(){
			tsk.context.SetDone(true)
			t.wait.Add(1)
			go func(tk task){
				defer t.wait.Done()
				tsk.context.Now=lib.Time()
				if tk.filter!=nil{
					if tk.filter(tk.context){
						tk.do(tk.context)
					}
				}else{
					tk.do(tk.context)
				}
				tk.context.SetDone(false)
			}(tsk)
		}
	}
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
//RegisterTask
func RegisterTask(name string,taskNums int,call DoFunc){
	globalTaskManager.RegisterTask(name,taskNums,call)
}

//RegisterTaskByMethod
func RegisterTaskByMethod(name string,taskNums int,call func(*context.Context)){
	globalTaskManager.RegisterTask(name,taskNums,call)
}
//RegisterTaskByMethodAndFilter
func RegisterTaskByMethodAndFilter(name string,taskNums int,call func(*context.Context),filter func(*context.Context)bool){
	globalTaskManager.RegisterTaskByFilter(name,taskNums,call,filter)
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