package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//DeleteAddressResult
type DeleteAddressResult struct{
	commonentity.Result
}

//String
func(d DeleteAddressResult)String()string{
	return lib.ObjectToString(d)
}
