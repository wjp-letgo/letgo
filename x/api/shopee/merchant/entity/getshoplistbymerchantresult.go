package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetShopListByMerchantResult
type GetShopListByMerchantResult struct{
	commonentity.Result
	ShopList []ShopEntity `json:"shop_list"`
	IsCnsc bool `json:"is_cnsc"`
	More bool `json:"more"`
}

//String
func(g GetShopListByMerchantResult)String()string{
	return lib.ObjectToString(g)
}