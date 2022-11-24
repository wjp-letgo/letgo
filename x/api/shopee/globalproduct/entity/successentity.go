package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//GetPublishTaskResultSuccessEntity
type GetPublishTaskResultSuccessEntity struct{
	Region string `json:"region"`
	ShopID int64 `json:"shop_id"`
	ItemID int64 `json:"item_id"`
}

//String
func(p GetPublishTaskResultSuccessEntity)String()string{
	return lib.ObjectToString(p)
}