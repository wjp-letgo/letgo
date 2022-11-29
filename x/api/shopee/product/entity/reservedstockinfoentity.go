package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ReservedStockInfoEntity
type ReservedStockInfoEntity struct {
	StockType       int    `json:"stock_type"`
	StockLocationID string `json:"stock_location_id"`
	ReservedStock   int    `json:"reserved_stock"`
}

//String
func (p ReservedStockInfoEntity) String() string {
	return lib.ObjectToString(p)
}
