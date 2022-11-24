package limiting

//Ilimiter
type Ilimiter interface{
	//是否达到限流 true 达到限流条件 false 还未达到限流
	Pass() bool
}