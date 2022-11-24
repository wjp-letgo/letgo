package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//UpdateSizeChartResult
type UpdateSizeChartResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(r UpdateSizeChartResult)String()string{
	return lib.ObjectToString(r)
}