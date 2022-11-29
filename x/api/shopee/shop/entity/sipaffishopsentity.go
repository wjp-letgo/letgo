package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//SipAffiShopsEntity
type SipAffiShopsEntity struct{
	AffiShopID int64 `json:"affi_shop_id"`
	Region string `json:"region"`
}

//String
func(g SipAffiShopsEntity)String()string{
	return lib.ObjectToString(g)
}