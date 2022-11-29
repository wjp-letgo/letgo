package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//TierVariationOptionLengthLimitEntity
type TierVariationOptionLengthLimitEntity struct {
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func (p TierVariationOptionLengthLimitEntity) String() string {
	return lib.ObjectToString(p)
}
