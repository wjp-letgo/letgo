package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetMerchantInfoResult
type GetMerchantInfoResult struct {
	commonentity.Result
	MerchantName string `json:"merchant_name"`
	IsCnsc       bool   `json:"is_cnsc"`
	AuthTime     int    `json:"auth_time"`
	ExpireTime   int    `json:"expire_time"`
}

//String
func (g GetMerchantInfoResult) String() string {
	return lib.ObjectToString(g)
}
