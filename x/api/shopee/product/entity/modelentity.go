package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ModelEntity
type ModelEntity struct {
	PriceInfo   []PriceInfoEntity `json:"price_info"`
	ModelID     int64             `json:"model_id"`
	StockInfo   []StockInfoEntity `json:"stock_info"`
	TierIndex   []int             `json:"tier_index"`
	PromotionID int64             `json:"promotion_id"`
	ModelSku    string            `json:"model_sku"`
}

//String
func (s ModelEntity) String() string {
	return lib.ObjectToString(s)
}

//InitTierVariationModelEntity
type InitTierVariationModelEntity struct {
	TierIndex     []int   `json:"tier_index"`
	NormalStock   int     `json:"normal_stock"`
	OriginalPrice float32 `json:"original_price"`
	ModelSku      string  `json:"model_sku"`
}

//String
func (s InitTierVariationModelEntity) String() string {
	return lib.ObjectToString(s)
}

//UpdateModelEntity
type UpdateModelEntity struct {
	ModelID  int64  `json:"model_id"`
	ModelSku string `json:"model_sku"`
}

//String
func (s UpdateModelEntity) String() string {
	return lib.ObjectToString(s)
}
