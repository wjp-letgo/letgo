package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//RecipientAddress
type RecipientAddressEntity struct{
	Name string `json:"name"`
	Phone string `json:"phone"`
	Town string `json:"town"`
	District string `json:"district"`
	City string `json:"city"`
	State string `json:"state"`
	Region string `json:"region"`
	Zipcode string `json:"zipcode"`
	FullAddress string `json:"full_address"`
}

//String
func(r RecipientAddressEntity)String()string{
	return lib.ObjectToString(r)
}