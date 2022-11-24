package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//BatchShipOrderResult
type BatchShipOrderResult struct{
	commonentity.Result
	Warning []ShippingDocumentParameterRequestOrderListEntity `json:"warning"`
	Response BatchShipOrderResultResponse `json:"response"`
}

//String
func(b BatchShipOrderResult)String()string{
	return lib.ObjectToString(b)
}

//BatchShipOrderResultResponse
type BatchShipOrderResultResponse struct{
	ResultList []BatchShipOrderResultEntity `json:"result_list"`
}

//BatchShipOrderResultEntity
type BatchShipOrderResultEntity struct{
	OrderSn string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
	FailError string `json:"fail_error"`
	FailMessage string `json:"fail_message"`
}

//String
func(b BatchShipOrderResultEntity)String()string{
	return lib.ObjectToString(b)
}