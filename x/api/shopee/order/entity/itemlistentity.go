package entity

import (
	"github.com/wjpxxx/letgo/lib"
)
//ItemListEntity
type ItemListEntity struct{
	ItemID int64 `json:"item_id"`
	OrderItemID int64 `json:"order_item_id"`
	ItemName string `json:"item_name"`
	ItemSku string `json:"item_sku"`
	ModelID int64 `json:"model_id"`
	ModelName string `json:"model_name"`
	ModelSku string `json:"model_sku"`
	ModelQuantityPurchased int64 `json:"model_quantity_purchased"`
	ModelOriginalPrice float32 `json:"model_original_price"`
	ModelDiscountedPrice float32 `json:"model_discounted_price"`
	Wholesale bool `json:"wholesale"`
	Weight float32 `json:"weight"`
	AddOnDeal bool `json:"add_on_deal"`
	MainItem bool `json:"main_item"`
	AddOnDealID int64 `json:"add_on_deal_id"`
	PromotionType string `json:"promotion_type"`
	PromotionID int64 `json:"promotion_id"`
	PromotionGroupID int64 `json:"promotion_group_id"`
}

//String
func(i ItemListEntity)String()string{
	return lib.ObjectToString(i)
}


//ItemListEntity
type PackageListRequestItemListEntity struct{
	ItemID int64 `json:"item_id"`
	ModelID int64 `json:"model_id"`
	OrderItemID int64 `json:"order_item_id"`
	PromotionGroupID int64 `json:"promotion_group_id"`
}

//String
func(p PackageListRequestItemListEntity)String()string{
	return lib.ObjectToString(p)
}