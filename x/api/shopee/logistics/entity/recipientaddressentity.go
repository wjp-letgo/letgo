package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//RecipientAddressEntity
type RecipientAddressEntity struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Town        string `json:"town"`
	District    string `json:"district"`
	City        string `json:"city"`
	State       string `json:"state"`
	Region      string `json:"region"`
	ZipCode     string `json:"zipcode"`
	FullAddress string `json:"full_address"`
}

//String
func (r RecipientAddressEntity) String() string {
	return lib.ObjectToString(r)
}
