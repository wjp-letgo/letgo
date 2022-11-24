package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//UpdateSipItemPriceResult
type UpdateSipItemPriceResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(r UpdateSipItemPriceResult)String()string{
	return lib.ObjectToString(r)
}