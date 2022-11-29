package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//GlobalItemEntity
type GlobalItemEntity struct {
	GlobalItemID     int64             `json:"global_item_id"`
	GlobalItemName   string            `json:"global_item_name"`
	Description      string            `json:"description"`
	GlobalItemSku    string            `json:"global_item_sku"`
	GlobalItemStatus string            `json:"global_item_status"`
	CreateTime       int               `json:"create_time"`
	UpdateTime       int               `json:"update_time"`
	StockInfo        StockInfoEntity   `json:"stock_info"`
	PriceInfo        PriceInfoEntity   `json:"price_info"`
	Image            ImageEntity       `json:"image"`
	Weight           float32           `json:"weight"`
	Dimension        DimensionEntity   `json:"dimension"`
	PreOrder         PreOrderEntity    `json:"pre_order"`
	SizeChart        string            `json:"size_chart"`
	Condition        string            `json:"condition"`
	HasModel         bool              `json:"has_model"`
	Video            VideoEntity       `json:"video"`
	CategoryID       int64             `json:"category_id"`
	Brand            BrandEntity       `json:"brand"`
	AttributeList    []AttributeEntity `json:"attribute_list"`
}

//String
func (c GlobalItemEntity) String() string {
	return lib.ObjectToString(c)
}
