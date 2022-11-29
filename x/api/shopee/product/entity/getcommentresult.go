package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetCommentResult
type GetCommentResult struct {
	commonentity.Result
	Response GetCommentResultResponse `json:"response"`
}

//String
func (g GetCommentResult) String() string {
	return lib.ObjectToString(g)
}

//GetCommentResultResponse
type GetCommentResultResponse struct {
	More            bool                `json:"more"`
	ItemCommentList []ItemCommentEntity `json:"item_comment_list"`
	NextCursor      string              `json:"next_cursor"`
}

//String
func (g GetCommentResultResponse) String() string {
	return lib.ObjectToString(g)
}
