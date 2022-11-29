package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//CancelOrderResult
type CancelOrderResult struct{
	commonentity.Result
	Response OrderUpdateTimeResponse `json:"response"`
}

//String
func(c CancelOrderResult)String()string{
	return lib.ObjectToString(c)
}

type OrderUpdateTimeResponse struct{
	UpdateTime int `json:"update_time"`
}

//String
func(o OrderUpdateTimeResponse)String()string{
	return lib.ObjectToString(o)
}

//CancelOrderRequestEntity
type CancelOrderRequestEntity struct{
	ItemID int64 `json:"item_id"`
	ModelID int64 `json:"model_id"`
}

//String
func(c CancelOrderRequestEntity)String()string{
	return lib.ObjectToString(c)
}