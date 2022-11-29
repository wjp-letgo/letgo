package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetShopListByMerchantResult
type GetShopListByMerchantResult struct {
	commonentity.Result
	ShopList []ShopEntity `json:"shop_list"`
	IsCnsc   bool         `json:"is_cnsc"`
	More     bool         `json:"more"`
}

//String
func (g GetShopListByMerchantResult) String() string {
	return lib.ObjectToString(g)
}
