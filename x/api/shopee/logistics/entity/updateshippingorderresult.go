package entity

import(
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/lib"
)

//UpdateShippingOrderResult
type UpdateShippingOrderResult struct{
	commonentity.Result
}

//String
func(u UpdateShippingOrderResult)String()string{
	return lib.ObjectToString(u)
}