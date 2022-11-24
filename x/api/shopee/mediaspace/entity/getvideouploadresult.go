package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetVideoUploadResult
type GetVideoUploadResult struct{
	commonentity.Result
	Error string `json:"error"`
	Warning string `json:"warning"`
	Response  GetVideoUploadResultResponse `json:"response"`
}

//String
func(g GetVideoUploadResult)String()string{
	return lib.ObjectToString(g)
}
//GetVideoUploadResultResponse
type GetVideoUploadResultResponse struct{
	Status string `json:"status"`
	VideoInfo VideoInfoEntity `json:"video_info"`
	Message string `json:"message"`
}

//String
func(g GetVideoUploadResultResponse)String()string{
	return lib.ObjectToString(g)
}