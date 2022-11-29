package limiting

import (
	"github.com/wjp-letgo/letgo/lib"
)

//漏桶算法
type LeakyBucket struct{
	counter float32 //当前请求数量
	counterLimit int //桶的大小
	rate float32  //桶漏的速度  每秒漏出几个请求  请求/ms
	lastTime int64  //最后一次请求的时间
}

//Pass 是否达到限流 true 达到限流条件 false 还未达到限流
func (l *LeakyBucket)Pass() bool {
	now:=lib.TimeLong()
	l.flowCounter(now)
	l.addCounter(now)
	r:= lib.Round(l.counter)>l.counterLimit
	if !r{
		l.lastTime=now
	}
	return r
}
//addCounter 增加水
func (l *LeakyBucket)addCounter(now int64){
	if lib.Round(l.counter)+1<=l.counterLimit+1{
		l.counter+=1
	}
}
//flowCounter 流出水
func (l *LeakyBucket)flowCounter(now int64){
	dis:=now-l.lastTime
	ot:=l.counter-(l.rate*float32(dis))
	if ot<0{
		l.counter=0
	}else{
		l.counter=ot
	}
}

//NewLeakyBucket
func NewLeakyBucket(counterLimit int, rate float32)*LeakyBucket{
	if rate<0.001{
		rate=0.001
	}
	return &LeakyBucket{0,counterLimit,rate,lib.TimeLong()}
}

//NewLeakyBucketByParams
func NewLeakyBucketByParams(args ...interface{})*LeakyBucket{
	if len(args)==1{
		return NewLeakyBucket(lib.InterfaceToInt(args[0]), 0.2)
	}else if len(args)==2{
		return NewLeakyBucket(lib.InterfaceToInt(args[0]),(&lib.Data{Value:args[1]}).Float32())
	}else{
		return NewLeakyBucket(200, 0.2)
	}
}