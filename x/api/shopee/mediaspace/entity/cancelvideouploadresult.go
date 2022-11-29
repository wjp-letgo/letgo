package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//CancelVideoUploadResult
type CancelVideoUploadResult struct {
	commonentity.Result
	Error   string `json:"error"`
	Warning string `json:"warning"`
}

//String
func (g CancelVideoUploadResult) String() string {
	return lib.ObjectToString(g)
}
