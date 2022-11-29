package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//InitTierVariationResult
type InitTierVariationResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g InitTierVariationResult) String() string {
	return lib.ObjectToString(g)
}
