package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ModelEntity
type ModelEntity struct {
	PriceInfo      []PriceInfoEntity `json:"price_info"`
	GlobalModelID  int64             `json:"global_model_id"`
	StockInfo      []StockInfoEntity `json:"stock_info"`
	TierIndex      []int             `json:"tier_index"`
	PromotionID    int64             `json:"promotion_id"`
	GlobalModelSku string            `json:"global_model_sku"`
}

//String
func (s ModelEntity) String() string {
	return lib.ObjectToString(s)
}

//InitTierVariationModelEntity
type InitTierVariationModelEntity struct {
	TierIndex      []int   `json:"tier_index"`
	NormalStock    int     `json:"normal_stock"`
	OriginalPrice  float32 `json:"original_price"`
	GlobalModelSku string  `json:"global_model_sku"`
}

//String
func (s InitTierVariationModelEntity) String() string {
	return lib.ObjectToString(s)
}

//UpdateModelEntity
type UpdateModelEntity struct {
	GlobalModelID  int64  `json:"global_model_id"`
	GlobalModelSku string `json:"global_model_sku"`
}

//String
func (s UpdateModelEntity) String() string {
	return lib.ObjectToString(s)
}

//CreatePublishTaskModelEntity
type CreatePublishTaskModelEntity struct {
	TierIndex     []int   `json:"tier_index"`
	OriginalPrice float32 `json:"original_price"`
}

//String
func (s CreatePublishTaskModelEntity) String() string {
	return lib.ObjectToString(s)
}
