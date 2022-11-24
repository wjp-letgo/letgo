package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//UpdateGlobalSizeChartResult
type UpdateGlobalSizeChartResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(r UpdateGlobalSizeChartResult)String()string{
	return lib.ObjectToString(r)
}