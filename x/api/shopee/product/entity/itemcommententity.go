package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ItemCommentEntity
type ItemCommentEntity struct {
	OrderSn       string             `json:"order_sn"`
	CommentID     int64              `json:"comment_id"`
	Comment       string             `json:"comment"`
	BuyerUsername string             `json:"buyer_username"`
	ItemID        int64              `json:"item_id"`
	ModelID       int64              `json:"model_id"`
	RatingStar    int                `json:"rating_star"`
	Editable      string             `json:"editable"`
	Hidden        bool               `json:"hidden"`
	CreateTime    int                `json:"create_time"`
	CommentReply  CommentReplyEntity `json:"comment_reply"`
}

//String
func (i ItemCommentEntity) String() string {
	return lib.ObjectToString(i)
}
