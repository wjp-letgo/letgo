package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//UnlistItemResult
type UnlistItemResult struct {
	commonentity.Result
	Warning  string                   `json:"warning"`
	Response UnlistItemResultResponse `json:"response"`
}

//String
func (r UnlistItemResult) String() string {
	return lib.ObjectToString(r)
}

//UnlistItemResultResponse
type UnlistItemResultResponse struct {
	FailureList []FailureEntity `json:"failure_list"`
	SuccessList []SuccessEntity `json:"success_list"`
}

//String
func (r UnlistItemResultResponse) String() string {
	return lib.ObjectToString(r)
}
