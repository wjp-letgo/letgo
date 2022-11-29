package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//CommentReplyEntity
type CommentReplyEntity struct {
	Reply  string `json:"reply"`
	Hidden bool   `json:"hidden"`
}

//String
func (c CommentReplyEntity) String() string {
	return lib.ObjectToString(c)
}
