package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//CommentReplyEntity
type CommentReplyEntity struct{
	Reply string `json:"reply"`
	Hidden bool `json:"hidden"`
}

//String
func(c CommentReplyEntity)String()string{
	return lib.ObjectToString(c)
}