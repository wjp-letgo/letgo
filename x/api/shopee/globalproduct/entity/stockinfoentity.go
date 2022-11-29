package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//StockInfoEntity
type StockInfoEntity struct {
	StockType       int    `json:"stock_type"`
	StockLocationID string `json:"stock_location_id"`
	NormalStock     int    `json:"normal_stock"`
	ReservedStock   int    `json:"reserved_stock"`
}

//String
func (p StockInfoEntity) String() string {
	return lib.ObjectToString(p)
}

//UpdateStockStockInfoEntity
type UpdateStockStockInfoEntity struct {
	GlobalModelID int64 `json:"global_model_id"`
	NormalStock   int   `json:"normal_stock"`
}

//String
func (p UpdateStockStockInfoEntity) String() string {
	return lib.ObjectToString(p)
}
