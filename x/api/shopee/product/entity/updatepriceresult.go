package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//UpdatePriceResult
type UpdatePriceResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response UpdatePriceResultResponse `json:"response"`
}

//String
func(g UpdatePriceResult)String()string{
	return lib.ObjectToString(g)
}

//UpdatePriceResultResponse
type UpdatePriceResultResponse struct{
	FailureList []UpdatePriceFailureEntity `json:"failure_list"`
	SuccessList  []UpdatePriceSuccessEntity `json:"success_list"`
}

//String
func(g UpdatePriceResultResponse)String()string{
	return lib.ObjectToString(g)
}