package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//InitVideoUploadResult
type InitVideoUploadResult struct {
	commonentity.Result
	Error    string                        `json:"error"`
	Response InitVideoUploadResultResponse `json:"response"`
}

//String
func (g InitVideoUploadResult) String() string {
	return lib.ObjectToString(g)
}

//InitVideoUploadResultResponse
type InitVideoUploadResultResponse struct {
	VideoUploadID string `json:"video_upload_id"`
}

//String
func (g InitVideoUploadResultResponse) String() string {
	return lib.ObjectToString(g)
}
