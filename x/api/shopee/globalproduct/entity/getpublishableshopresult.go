package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetPublishableShopResult
type GetPublishableShopResult struct{
	commonentity.Result
	Response GetPublishableShopResultResponse `json:"response"`
	Warning string `json:"warning"`
}

//String
func(g GetPublishableShopResult)String()string{
	return lib.ObjectToString(g)
}
//GetPublishableShopResultResponse
type GetPublishableShopResultResponse struct{
	PublishableShop []PublishableShopEntity `json:"publishable_shop"`
}

//String
func(g GetPublishableShopResultResponse)String()string{
	return lib.ObjectToString(g)
}