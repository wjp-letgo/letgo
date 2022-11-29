package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//PromotionPriceInfoEntity
type PromotionPriceInfoEntity struct {
	PromotionPrice float32 `json:"promotion_price"`
}

//String
func (p PromotionPriceInfoEntity) String() string {
	return lib.ObjectToString(p)
}
