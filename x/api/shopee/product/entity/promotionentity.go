package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//PromotionEntity
type PromotionEntity struct {
	PromotionType      string                     `json:"promotion_type"`
	PromotionID        int64                      `json:"promotion_id"`
	ModelID            int64                      `json:"model_id"`
	StartTime          int                        `json:"start_time"`
	EndTime            int                        `json:"end_time"`
	PromotionPriceInfo []PromotionPriceInfoEntity `json:"promotion_price_info"`
	ReservedStockInfo  []ReservedStockInfoEntity  `json:"reserved_stock_info"`
	PromotionStaging   string                     `json:"promotion_staging"`
}

//String
func (p PromotionEntity) String() string {
	return lib.ObjectToString(p)
}
