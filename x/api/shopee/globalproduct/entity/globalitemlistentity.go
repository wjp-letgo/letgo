package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//GlobalItemListEntity
type GlobalItemListEntity struct {
	GlobalItemID int64 `json:"global_item_id"`
	UpdateTime   int   `json:"update_time"`
}

//String
func (i GlobalItemListEntity) String() string {
	return lib.ObjectToString(i)
}
