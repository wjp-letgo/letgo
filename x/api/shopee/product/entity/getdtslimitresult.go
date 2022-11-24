package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetDtsLimitResult
type GetDtsLimitResult struct{
	commonentity.Result
	Response GetDtsLimitResultResponse `json:"response"`
	Warning string `json:"warning"`
}

//String
func(g GetDtsLimitResult)String()string{
	return lib.ObjectToString(g)
}
//GetDtsLimitResultResponse
type GetDtsLimitResultResponse struct{
	DaysToShipLimit DaysToShipLimitEntity `json:"days_to_ship_limit"`
	NonPreOrderDaysToShip int `json:"non_pre_order_days_to_ship"`
}

//String
func(g GetDtsLimitResultResponse)String()string{
	return lib.ObjectToString(g)
}