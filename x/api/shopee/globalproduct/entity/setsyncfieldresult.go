package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//SetSyncFieldResult
type SetSyncFieldResult struct{
	commonentity.Result
	Warning string `json:"warning"`
}

//String
func(r SetSyncFieldResult)String()string{
	return lib.ObjectToString(r)
}
