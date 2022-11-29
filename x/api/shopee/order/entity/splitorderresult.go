package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//SplitOrderResult
type SplitOrderResult struct{
	commonentity.Result
	Response SplitOrderResponse `json:"response"`
}

//String
func(s SplitOrderResult)String()string{
	return lib.ObjectToString(s)
}
//SplitOrderResponse
type SplitOrderResponse struct{
	OrderSn string `json:"order_sn"`
	PackageList []PackageListEntity `json:"package_list"`
}

//String
func(s SplitOrderResponse)String()string{
	return lib.ObjectToString(s)
}