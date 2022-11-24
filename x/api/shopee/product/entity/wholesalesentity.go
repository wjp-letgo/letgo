package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//WholesalesEntity
type WholesalesEntity struct{
	MinCount int `json:"min_count"`
	MaxCount int `json:"max_count"`
	UnitPrice float32 `json:"unit_price"`
	InflatedPriceOfUnitPrice float32 `json:"inflated_price_of_unit_price"`
}

//String
func(w WholesalesEntity)String()string{
	return lib.ObjectToString(w)
}