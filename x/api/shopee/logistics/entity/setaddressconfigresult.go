package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//AddressTypeConfigEntity
type SetAddressConfigResult struct{
	commonentity.Result
}

//String
func(a SetAddressConfigResult)String()string{
	return lib.ObjectToString(a)
}
