package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//ReplyCommentRequestCommentEntity
type ReplyCommentRequestCommentEntity struct{
	CommentID int64 `json:"comment_id"`
	Comment string `json:"comment"`
}

//String
func(c ReplyCommentRequestCommentEntity)String()string{
	return lib.ObjectToString(c)
}