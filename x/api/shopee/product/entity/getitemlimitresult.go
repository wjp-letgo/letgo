package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetItemLimitResult
type GetItemLimitResult struct {
	commonentity.Result
	Response GetItemLimitResultResponse `json:"response"`
	Warning  string                     `json:"warning"`
}

//String
func (g GetItemLimitResult) String() string {
	return lib.ObjectToString(g)
}

//GetItemLimitResultResponse
type GetItemLimitResultResponse struct {
	PriceLimit                        PriceLimitEntity                        `json:"price_limit"`
	WholesalePriceThresholdPercentage WholesalePriceThresholdPercentageEntity `json:"wholesale_price_threshold_percentage"`
	StockLimit                        StockLimitEntity                        `json:"stock_limit"`
	ItemNameLengthLimit               ItemNameLengthLimitEntity               `json:"item_name_length_limit"`
	ItemImageCountLimit               ItemImageCountLimitEntity               `json:"item_image_count_limit"`
	ItemDescriptionLengthLimit        ItemDescriptionLengthLimitEntity        `json:"item_description_length_limit"`
	TierVariationNameLengthLimit      TierVariationNameLengthLimitEntity      `json:"tier_variation_name_length_limit"`
	TierVariationOptionLengthLimit    TierVariationOptionLengthLimitEntity    `json:"tier_variation_option_length_limit"`
	ItemCountLimit                    ItemCountLimitEntity                    `json:"item_count_limit"`
}

//String
func (g GetItemLimitResultResponse) String() string {
	return lib.ObjectToString(g)
}
