package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//LogisticEntity
type LogisticEntity struct{
	LogisticID int64 `json:"logistic_id"`
	LogisticName string `json:"logistic_name"`
	Enabled bool `json:"enabled"`
	ShippingFee float32 `json:"shipping_fee"`
	SizeID int64 `json:"size_id"`
	IsFree bool `json:"is_free"`
	EstimatedShippingFee bool `json:"estimated_shipping_fee"`
}

//String
func(l LogisticEntity)String()string{
	return lib.ObjectToString(l)
}