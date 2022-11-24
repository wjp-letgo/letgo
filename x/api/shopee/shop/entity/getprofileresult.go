package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetProfileResult
type GetProfileResult struct{
	commonentity.Result
	Response GetProfileResultResponse `json:"response"`
}

//String
func(g GetProfileResult)String()string{
	return lib.ObjectToString(g)
}
//GetProfileResultResponse
type GetProfileResultResponse struct{
	ShopLogo string `json:"shop_logo"`
	Description string `json:"description"`
	ShopName string `json:"shop_name"`
}
//String
func(g GetProfileResultResponse)String()string{
	return lib.ObjectToString(g)
}