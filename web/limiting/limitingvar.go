package limiting

//LimitFlowType 限流算法
type LimitFlowType string

const (
	LIMIT_FLOW_COUNTER LimitFlowType ="Counter"            	               //计数器算法-固定窗口
	LIMIT_FLOW_ROLLING_COUNTER LimitFlowType ="Rolling counter"            //计数器算法-滑动窗口
	LIMIT_FLOW_LEAKY_BUCKET LimitFlowType ="Leaky bucket"  //漏桶算法
	LIMIT_FLOW_TOKEN_BUCKET LimitFlowType ="Token bucket"  //令牌桶
	LIMIT_FLOW_NONE LimitFlowType ="None"  //不限流
)