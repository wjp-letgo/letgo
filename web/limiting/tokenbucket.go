package limiting
import (
	"github.com/wjpxxx/letgo/lib"
)

//令牌桶算法
type TokenBucket struct{
	counter float32 //当前请求数量
	rate float32  //进水的速度  每秒漏出几个请求  请求/ms
	lastTime int64  //最后一次请求的时间
}

//Pass 是否达到限流 true 达到限流条件 false 还未达到限流
func (l *TokenBucket)Pass() bool {
	now:=lib.TimeLong()
	l.flowCounter(now)
	l.subCounter(now)
	r:= lib.Round(l.counter)<=0
	if !r{
		l.lastTime=now
	}
	return r
}
//subCounter 取走token
func (l *TokenBucket)subCounter(now int64){
	s:=l.counter-1
	if s<= -1{
	}else{
		l.counter=s
	}
}
//flowCounter 水流进的速度
func (l *TokenBucket)flowCounter(now int64){
	dis:=now-l.lastTime
	l.counter=l.counter+l.rate*float32(dis)
}

//NewTokenBucket
func NewTokenBucket(rate float32)*TokenBucket{
	return &TokenBucket{0,rate,lib.TimeLong()}
}

//NewTokenBucketByParams
func NewTokenBucketByParams(args ...interface{})*TokenBucket{
	if len(args)==1{
		return NewTokenBucket((&lib.Data{Value:args[0]}).Float32())
	}else{
		return NewTokenBucket(0.2)
	}
}