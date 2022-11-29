package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetTrackingInfoResult
type GetTrackingInfoResult struct{
	commonentity.Result
	Response GetTrackingInfoResponse `json:"response"`
}

//String
func(g GetTrackingInfoResult)String()string{
	return lib.ObjectToString(g)
}

//GetShippingParameterResponse
type GetTrackingInfoResponse struct{
	OrderSn string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
	LogisticsStatus string `json:"logistics_status"`
	TrackingInfo []TrackingInfoEntity `json:"tracking_info"`
}

//String
func(g GetTrackingInfoResponse)String()string{
	return lib.ObjectToString(g)
}