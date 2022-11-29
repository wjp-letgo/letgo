package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetShippingDocumentParameterResult
type GetShippingDocumentParameterResult struct {
	commonentity.Result
	Warning  []ShippingDocumentParameterRequestOrderListEntity `json:"warning"`
	Response GetShippingDocumentParameterResultResponse        `json:"response"`
}

//String
func (g GetShippingDocumentParameterResult) String() string {
	return lib.ObjectToString(g)
}

//GetShippingDocumentParameterResultResponse
type GetShippingDocumentParameterResultResponse struct {
	ResultList []GetShippingDocumentParameterResultEntity `json:"result_list"`
}

//String
func (g GetShippingDocumentParameterResultResponse) String() string {
	return lib.ObjectToString(g)
}

//GetShippingDocumentParameterResultEntity
type GetShippingDocumentParameterResultEntity struct {
	OrderSn                        string   `json:"order_sn"`
	PackageNumber                  string   `json:"package_number"`
	SuggestShippingDocumentType    string   `json:"suggest_shipping_document_type"`
	SelectableShippingDocumentType []string `json:"selectable_shipping_document_type"`
	FailError                      string   `json:"fail_error"`
	FailMessage                    string   `json:"fail_message"`
}

//String
func (g GetShippingDocumentParameterResultEntity) String() string {
	return lib.ObjectToString(g)
}
