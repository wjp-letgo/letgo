package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//SetSyncFieldResult
type SetSyncFieldResult struct {
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func (r SetSyncFieldResult) String() string {
	return lib.ObjectToString(r)
}
