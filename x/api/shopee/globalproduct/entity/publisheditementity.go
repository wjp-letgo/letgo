package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//PublishedItemEntity
type PublishedItemEntity struct {
	ShopID     int64  `json:"shop_id"`
	ShopRegion string `json:"shop_region"`
	ItemID     int64  `json:"item_id"`
	ItemStatus int    `json:"item_status"`
}

//String
func (p PublishedItemEntity) String() string {
	return lib.ObjectToString(p)
}
