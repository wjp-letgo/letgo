package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetPublishTaskResult
type GetPublishTaskResult struct{
	commonentity.Result
	Response GetPublishTaskResultResponse `json:"response"`
	Warning string `json:"warning"`
}

//String
func(g GetPublishTaskResult)String()string{
	return lib.ObjectToString(g)
}
//GetPublishTaskResultResponse
type GetPublishTaskResultResponse struct{
	PublishStatus string `json:"publish_status"`
	Success GetPublishTaskResultSuccessEntity `json:"success"`
	Failed GetPublishTaskResultFailureEntity `json:"failed"`
}

//String
func(g GetPublishTaskResultResponse)String()string{
	return lib.ObjectToString(g)
}