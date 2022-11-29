package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetShippingDocumentInfoResult
type GetShippingDocumentInfoResult struct {
	commonentity.Result
	Response GetShippingDocumentInfoResponse `json:"response"`
}

//String
func (g GetShippingDocumentInfoResult) String() string {
	return lib.ObjectToString(g)
}

//GetShippingDocumentInfoResponse
type GetShippingDocumentInfoResponse struct {
	ShippingDocumentInfo ShippingDocumentInfoEntity `json:"shipping_document_info"`
}

//String
func (g GetShippingDocumentInfoResponse) String() string {
	return lib.ObjectToString(g)
}
