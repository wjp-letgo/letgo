package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//CreateShippingDocumentResult
type CreateShippingDocumentResult struct {
	commonentity.Result
	Warning  []ShippingDocumentParameterRequestOrderListEntity `json:"warning"`
	Response CreateShippingDocumentResultResponse              `json:"response"`
}

//String
func (c CreateShippingDocumentResult) String() string {
	return lib.ObjectToString(c)
}

//CreateShippingDocumentResultResponse
type CreateShippingDocumentResultResponse struct {
	ResultList []CreateShippingDocumentResultEntity `json:"result_list"`
}

//String
func (c CreateShippingDocumentResultResponse) String() string {
	return lib.ObjectToString(c)
}

//CreateShippingDocumentResultEntity
type CreateShippingDocumentResultEntity struct {
	OrderSn       string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
	FailError     string `json:"fail_error"`
	FailMessage   string `json:"fail_message"`
}

//String
func (c CreateShippingDocumentResultEntity) String() string {
	return lib.ObjectToString(c)
}
