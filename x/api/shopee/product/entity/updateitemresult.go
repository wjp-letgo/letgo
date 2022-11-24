package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)


//UpdateItemResult
type UpdateItemResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response ItemEntity `json:"response"`
}

//String
func(g UpdateItemResult)String()string{
	return lib.ObjectToString(g)
}