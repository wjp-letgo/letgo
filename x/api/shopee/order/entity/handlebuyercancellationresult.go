package entity

import(
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/lib"
)

//HandleBuyerCancellationResult
type HandleBuyerCancellationResult struct{
	commonentity.Result
}

//String
func(h HandleBuyerCancellationResult)String()string{
	return lib.ObjectToString(h)
}