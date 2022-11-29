package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetPublishTaskResult
type GetPublishTaskResult struct {
	commonentity.Result
	Response GetPublishTaskResultResponse `json:"response"`
	Warning  string                       `json:"warning"`
}

//String
func (g GetPublishTaskResult) String() string {
	return lib.ObjectToString(g)
}

//GetPublishTaskResultResponse
type GetPublishTaskResultResponse struct {
	PublishStatus string                            `json:"publish_status"`
	Success       GetPublishTaskResultSuccessEntity `json:"success"`
	Failed        GetPublishTaskResultFailureEntity `json:"failed"`
}

//String
func (g GetPublishTaskResultResponse) String() string {
	return lib.ObjectToString(g)
}
