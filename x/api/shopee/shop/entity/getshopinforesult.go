package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetShopInfoResult
type GetShopInfoResult struct {
	ShopName     string               `json:"shop_name"`
	Region       string               `json:"region"`
	Status       string               `json:"status"`
	SipAffiShops []SipAffiShopsEntity `json:"sip_affi_shops"`
	IsCb         bool                 `json:"is_cb"`
	IsCnsc       bool                 `json:"is_cnsc"`
	commonentity.Result
	AuthTime   int `json:"auth_time"`
	ExpireTime int `json:"expire_time"`
}

//String
func (g GetShopInfoResult) String() string {
	return lib.ObjectToString(g)
}
