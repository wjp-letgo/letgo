package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//FailureEntity
type FailureEntity struct{
	ItemID int64 `json:"item_id"`
	FailedReason string `json:"failed_reason"`
}

//String
func(d FailureEntity)String()string{
	return lib.ObjectToString(d)
}


//FailureEntity
type UpdatePriceFailureEntity struct{
	ModelID int64 `json:"model_id"`
	FailedReason string `json:"failed_reason"`
}

//String
func(d UpdatePriceFailureEntity)String()string{
	return lib.ObjectToString(d)
}


//UpdateStockFailureEntity
type UpdateStockFailureEntity struct{
	ModelID int64 `json:"model_id"`
	FailedReason string `json:"failed_reason"`
}

//String
func(d UpdateStockFailureEntity)String()string{
	return lib.ObjectToString(d)
}
//GetItemPromotionFailureEntity
type GetItemPromotionFailureEntity struct{
	ItemID int64 `json:"item_id"`
	FailedReason string `json:"failed_reason"`
}

//String
func(d GetItemPromotionFailureEntity)String()string{
	return lib.ObjectToString(d)
}