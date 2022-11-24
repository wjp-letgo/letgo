package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//PublishableShopEntity
type PublishableShopEntity struct{
	ShopID int64 `json:"shop_id"`
	ShopRegion string `json:"shop_region"`
}

//String
func(p PublishableShopEntity)String()string{
	return lib.ObjectToString(p)
}