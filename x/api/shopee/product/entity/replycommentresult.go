package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//ReplyCommentResult
type ReplyCommentResult struct{
	commonentity.Result
	Response ReplyCommentResultResponse `json:"response"`
	Warning []string `json:"warning"`
}

//String
func(r ReplyCommentResult)String()string{
	return lib.ObjectToString(r)
}

//ReplyCommentResultResponse
type ReplyCommentResultResponse struct{
	ResultList [] ReplyCommentResultList `json:"result_list"`
}

//String
func(r ReplyCommentResultResponse)String()string{
	return lib.ObjectToString(r)
}

//ReplyCommentResultList
type ReplyCommentResultList struct{
	CommentID int64 `json:"comment_id"`
	FailError string `json:"fail_error"`
	FailMessage string `json:"fail_message"`
}