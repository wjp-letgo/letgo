package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//PriceInfoEntity
type PriceInfoEntity struct{
	Currency string `json:"currency"`
	OriginalPrice float32 `json:"original_price"`
	CurrentPrice float32 `json:"current_price"`
	InflatedPriceOfOriginalPrice float32 `json:"inflated_price_of_original_price"`
	InflatedPriceOfCurrentPrice float32 `json:"inflated_price_of_current_price"`
	SipItemPrice float32 `json:"sip_item_price"`
	SipItemPriceSource string `json:"sip_item_price_source"`
}

//String
func(p PriceInfoEntity)String()string{
	return lib.ObjectToString(p)
}


//PriceInfoEntity
type UpdatePricePriceInfoEntity struct{
	ModelID int64 `json:"model_id"`
	OriginalPrice float32 `json:"original_price"`
}

//String
func(p UpdatePricePriceInfoEntity)String()string{
	return lib.ObjectToString(p)
}