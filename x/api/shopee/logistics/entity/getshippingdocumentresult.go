package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetShippingDocumentResult
type GetShippingDocumentResult struct {
	commonentity.Result
	Warning  []ShippingDocumentParameterRequestOrderListEntity `json:"warning"`
	Response GetShippingDocumentResultResponse                 `json:"response"`
}

//String
func (g GetShippingDocumentResult) String() string {
	return lib.ObjectToString(g)
}

//GetShippingDocumentResultResponse
type GetShippingDocumentResultResponse struct {
	ResultList []GetShippingDocumentResultEntity `json:"result_list"`
}

//String
func (g GetShippingDocumentResultResponse) String() string {
	return lib.ObjectToString(g)
}

//GetShippingDocumentResultEntity
type GetShippingDocumentResultEntity struct {
	OrderSn       string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
	Status        string `json:"status"`
	FailError     string `json:"fail_error"`
	FailMessage   string `json:"fail_message"`
}

//String
func (g GetShippingDocumentResultEntity) String() string {
	return lib.ObjectToString(g)
}
