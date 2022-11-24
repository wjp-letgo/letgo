package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetGlobalItemLimitResult
type GetGlobalItemLimitResult struct{
	commonentity.Result
	Response GetGlobalItemLimitResultResponse `json:"response"`
	Warning string `json:"warning"`
	TextLengthMultiplier float32 `json:"text_length_multiplier"`
}

//String
func(g GetGlobalItemLimitResult)String()string{
	return lib.ObjectToString(g)
}



//GetGlobalItemLimitResultResponse
type GetGlobalItemLimitResultResponse struct{
	PriceLimit PriceLimitEntity `json:"price_limit"`
	StockLimit StockLimitEntity `json:"stock_limit"`
	ItemNameLengthLimit ItemNameLengthLimitEntity `json:"global_item_name_length_limit"`
	ItemImageCountLimit ItemImageCountLimitEntity `json:"global_item_image_count_limit"`
	ItemDescriptionLengthLimit ItemDescriptionLengthLimitEntity `json:"global_item_description_length_limit"`
	TierVariationNameLengthLimit TierVariationNameLengthLimitEntity `json:"tier_variation_name_length_limit"`
	TierVariationOptionLengthLimit TierVariationOptionLengthLimitEntity `json:"tier_variation_option_length_limit"`
}

//String
func(g GetGlobalItemLimitResultResponse)String()string{
	return lib.ObjectToString(g)
}