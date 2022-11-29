package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//LogisticInfoEntity
type LogisticInfoEntity struct {
	LogisticID           int64   `json:"logistic_id"`
	LogisticName         string  `json:"logistic_name"`
	Enabled              bool    `json:"enabled"`
	ShippingFee          float32 `json:"shipping_fee"`
	SizeID               int64   `json:"size_id"`
	IsFree               bool    `json:"is_free"`
	EstimatedShippingFee bool    `json:"estimated_shipping_fee"`
}

//String
func (l LogisticInfoEntity) String() string {
	return lib.ObjectToString(l)
}
