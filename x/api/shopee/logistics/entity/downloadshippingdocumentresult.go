package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//DownloadShippingDocumentResult
type DownloadShippingDocumentResult struct{
	commonentity.Result
	File []byte
}

//String
func(d DownloadShippingDocumentResult)String()string{
	return lib.ObjectToString(d)
}