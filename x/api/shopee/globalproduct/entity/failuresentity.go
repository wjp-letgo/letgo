package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//FailureEntity
type FailureEntity struct {
	ShopID  int64 `json:"shop_id"`
	ItemID  int64 `json:"item_id"`
	ModelID int64 `json:"model_id"`
}

//String
func (p FailureEntity) String() string {
	return lib.ObjectToString(p)
}

//GetPublishTaskResultFailureEntity
type GetPublishTaskResultFailureEntity struct {
	FailedReason string `json:"failed_reason"`
}

//String
func (p GetPublishTaskResultFailureEntity) String() string {
	return lib.ObjectToString(p)
}
