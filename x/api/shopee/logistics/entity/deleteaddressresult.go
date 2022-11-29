package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//DeleteAddressResult
type DeleteAddressResult struct {
	commonentity.Result
}

//String
func (d DeleteAddressResult) String() string {
	return lib.ObjectToString(d)
}
