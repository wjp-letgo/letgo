package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//UpdateGlobalStockResult
type UpdateGlobalStockResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (r UpdateGlobalStockResult) String() string {
	return lib.ObjectToString(r)
}
