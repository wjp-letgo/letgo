package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//TierVariationOptionLengthLimitEntity
type TierVariationOptionLengthLimitEntity struct{
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func(p TierVariationOptionLengthLimitEntity)String()string{
	return lib.ObjectToString(p)
}