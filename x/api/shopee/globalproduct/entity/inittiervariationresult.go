package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//InitTierVariationResult
type InitTierVariationResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(g InitTierVariationResult)String()string{
	return lib.ObjectToString(g)
}