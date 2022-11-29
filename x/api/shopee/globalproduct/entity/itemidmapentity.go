package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ItemIdMapEntity
type ItemIdMapEntity struct {
	ItemID       int64 `json:"item_id"`
	GlobalItemID int64 `json:"global_item_id"`
}

//String
func (i ItemIdMapEntity) String() string {
	return lib.ObjectToString(i)
}
