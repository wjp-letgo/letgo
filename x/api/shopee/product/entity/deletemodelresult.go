package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//DeleteModelResult
type DeleteModelResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(g DeleteModelResult)String()string{
	return lib.ObjectToString(g)
}