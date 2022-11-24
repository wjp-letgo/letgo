package entity

import(
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/lib"
)

//UnSplitOrderResult
type AddInvoiceDataResult struct{
	commonentity.Result
}

//String
func(a AddInvoiceDataResult)String()string{
	return lib.ObjectToString(a)
}