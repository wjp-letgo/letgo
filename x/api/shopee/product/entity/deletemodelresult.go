package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//DeleteModelResult
type DeleteModelResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (g DeleteModelResult) String() string {
	return lib.ObjectToString(g)
}
