package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//FailureDeleteItemEntity
type FailureDeleteItemEntity struct{
	ShopID int64 `json:"shop_id"`
	ItemID int64 `json:"item_id"`
}

//String
func(p FailureDeleteItemEntity)String()string{
	return lib.ObjectToString(p)
}