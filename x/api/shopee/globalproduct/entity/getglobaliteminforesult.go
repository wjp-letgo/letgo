package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetGlobalItemInfoResult
type GetGlobalItemInfoResult struct{
	commonentity.Result
	Response GetGlobalItemInfoResultResponse `json:"response"`
	Warning string `json:"warning"`
}

//String
func(g GetGlobalItemInfoResult)String()string{
	return lib.ObjectToString(g)
}
//GetGlobalItemInfoResultResponse
type GetGlobalItemInfoResultResponse struct{
	GlobalItemList []GlobalItemEntity `json:"global_item_list"`
}
//String
func(g GetGlobalItemInfoResultResponse)String()string{
	return lib.ObjectToString(g)
}