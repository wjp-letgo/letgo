package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//SipItemPriceEntity
type SipItemPriceEntity struct {
	ModelID      int64   `json:"model_id"`
	SipItemPrice float32 `json:"sip_item_price"`
}

//String
func (p SipItemPriceEntity) String() string {
	return lib.ObjectToString(p)
}
