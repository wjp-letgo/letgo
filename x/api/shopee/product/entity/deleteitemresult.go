package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//DeleteItemResult
type DeleteItemResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(g DeleteItemResult)String()string{
	return lib.ObjectToString(g)
}