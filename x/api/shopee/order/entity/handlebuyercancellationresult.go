package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//HandleBuyerCancellationResult
type HandleBuyerCancellationResult struct{
	commonentity.Result
}

//String
func(h HandleBuyerCancellationResult)String()string{
	return lib.ObjectToString(h)
}