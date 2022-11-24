package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//ShopSyncEntity
type ShopSyncEntity struct{
	ShopID int64 `json:"shop_id"`
	ShopRegion string `json:"shop_region"`
	NameAndDescription bool `json:"name_and_description"`
	MediaInformation bool `json:"media_information"`
	Category bool `json:"category"`
	TierVariationNameAndOption bool `json:"tier_variation_name_and_option"`
	Price bool `json:"price"`
	DaysToShip bool `json:"days_to_ship"`
}

//String
func(p ShopSyncEntity)String()string{
	return lib.ObjectToString(p)
}