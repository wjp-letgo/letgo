package context

import(
	"sync"
)


//Context 任务上下文对象
type Context struct{
	Name string `json:"name"`
	TaskNo int `json:"taskNo"`
	Done bool `json:"done"`
	lock sync.RWMutex
	Now int `json:"now"`
}
//IsDone
func (c *Context)IsDone()bool{
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.Done
}
//SetDone
func (c *Context)SetDone(d bool){
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Done=d
}
//NewContext
func NewContext()*Context{
	return &Context{
		Done:false,
	}
}