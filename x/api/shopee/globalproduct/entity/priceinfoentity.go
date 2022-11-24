package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//PriceInfoEntity
type PriceInfoEntity struct{
	Currency string `json:"currency"`
	OriginalPrice float32 `json:"original_price"`
}

//String
func(p PriceInfoEntity)String()string{
	return lib.ObjectToString(p)
}
//PriceEntity
type PriceEntity struct{
	GlobalModelID int64 `json:"global_model_id"`
	OriginalPrice float32 `json:"original_price"`
}

//String
func(p PriceEntity)String()string{
	return lib.ObjectToString(p)
}