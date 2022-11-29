package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//UnSplitOrderResult
type UnSplitOrderResult struct{
	commonentity.Result
	Response OrderUpdateTimeResponse `json:"response"`
}

//String
func(u UnSplitOrderResult)String()string{
	return lib.ObjectToString(u)
}