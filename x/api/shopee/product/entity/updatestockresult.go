package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//UpdateStockResult
type UpdateStockResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response UpdateStockResultResponse `json:"response"`
}

//String
func(g UpdateStockResult)String()string{
	return lib.ObjectToString(g)
}

//UpdateStockResultResponse
type UpdateStockResultResponse struct{
	FailureList []UpdateStockFailureEntity `json:"failure_list"`
	SuccessList  []UpdateStockSuccessEntity `json:"success_list"`
}

//String
func(g UpdateStockResultResponse)String()string{
	return lib.ObjectToString(g)
}