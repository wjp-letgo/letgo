package entity

import(
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/lib"
)

//ShipOrderResult
type ShipOrderResult struct{
	commonentity.Result
}

//String
func(s ShipOrderResult)String()string{
	return lib.ObjectToString(s)
}