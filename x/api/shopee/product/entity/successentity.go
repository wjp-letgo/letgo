package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//SuccessEntity
type SuccessEntity struct{
	ItemID int64 `json:"item_id"`
	Unlist bool `json:"unlist"`
}

//String
func(d SuccessEntity)String()string{
	return lib.ObjectToString(d)
}

//BoostItemSuccessEntity
type BoostItemSuccessEntity struct{
	ItemIdList []int64 `json:"item_id_list"`
}

//String
func(d BoostItemSuccessEntity)String()string{
	return lib.ObjectToString(d)
}

//UpdatePriceSuccessEntity
type UpdatePriceSuccessEntity struct{
	ModelID int64 `json:"model_id"`
	OriginalPrice float32 `json:"original_price"`
}

//String
func(d UpdatePriceSuccessEntity)String()string{
	return lib.ObjectToString(d)
}


//UpdateStockSuccessEntity
type UpdateStockSuccessEntity struct{
	ModelID int64 `json:"model_id"`
	NormalStock int `json:"normal_stock"`
}

//String
func(d UpdateStockSuccessEntity)String()string{
	return lib.ObjectToString(d)
}
//GetItemPromotionSuccessEntity
type GetItemPromotionSuccessEntity struct{
	ItemID int64 `json:"item_id"`
	Promotion []PromotionEntity `json:"promotion"`
}

//String
func(d GetItemPromotionSuccessEntity)String()string{
	return lib.ObjectToString(d)
}