package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//CreatePublishTaskResult
type CreatePublishTaskResult struct {
	commonentity.Result
	Response CreatePublishTaskResultResponse `json:"response"`
	Warning  string                          `json:"warning"`
}

//String
func (g CreatePublishTaskResult) String() string {
	return lib.ObjectToString(g)
}

//CreatePublishTaskResultResponse
type CreatePublishTaskResultResponse struct {
	PublishTaskID int64 `json:"publish_task_id"`
}

//String
func (g CreatePublishTaskResultResponse) String() string {
	return lib.ObjectToString(g)
}
