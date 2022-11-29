package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//DeleteGlobalItemResult
type DeleteGlobalItemResult struct {
	commonentity.Result
	Response DeleteGlobalItemResultResponse `json:"response"`
	Warning  string                         `json:"warning"`
}

//String
func (g DeleteGlobalItemResult) String() string {
	return lib.ObjectToString(g)
}

//DeleteGlobalItemResultResponse
type DeleteGlobalItemResultResponse struct {
	FailureDeleteItem []FailureDeleteItemEntity `json:"failure_delete_item"`
}

//String
func (g DeleteGlobalItemResultResponse) String() string {
	return lib.ObjectToString(g)
}
