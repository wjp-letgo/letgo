package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetItemPromotionResult
type GetItemPromotionResult struct {
	commonentity.Result
	Response GetItemPromotionResultResponse `json:"response"`
	Warning  string                         `json:"warning"`
}

//String
func (g GetItemPromotionResult) String() string {
	return lib.ObjectToString(g)
}

//GetItemPromotionResultResponse
type GetItemPromotionResultResponse struct {
	SuccessList []GetItemPromotionSuccessEntity `json:"success_list"`
	FailureList []GetItemPromotionFailureEntity `json:"failure_list"`
}
