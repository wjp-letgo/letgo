package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetUnbindOrderListResult
type GetUnbindOrderListResult struct {
	commonentity.Result
	Warning  string                           `json:"warning"`
	Response GetUnbindOrderListResultResponse `json:"response"`
}

//String
func (g GetUnbindOrderListResult) String() string {
	return lib.ObjectToString(g)
}

//GetUnbindOrderListResultResponse
type GetUnbindOrderListResultResponse struct {
	More       bool          `json:"more"`
	NextCursor string        `json:"next_cursor"`
	OrderList  []OrderEntity `json:"order_list"`
}

//String
func (g GetUnbindOrderListResultResponse) String() string {
	return lib.ObjectToString(g)
}
