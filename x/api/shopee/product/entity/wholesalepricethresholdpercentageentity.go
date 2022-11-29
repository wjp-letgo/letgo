package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//PriceLimitEntity
type WholesalePriceThresholdPercentageEntity struct {
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func (p WholesalePriceThresholdPercentageEntity) String() string {
	return lib.ObjectToString(p)
}
