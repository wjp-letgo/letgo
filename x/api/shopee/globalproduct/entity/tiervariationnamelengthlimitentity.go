package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//TierVariationNameLengthLimitEntity
type TierVariationNameLengthLimitEntity struct{
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func(p TierVariationNameLengthLimitEntity)String()string{
	return lib.ObjectToString(p)
}