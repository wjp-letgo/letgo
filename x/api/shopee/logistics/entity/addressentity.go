package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//AddressEntity
type AddressEntity struct {
	AddressID    int64            `json:"address_id"`
	Region       string           `json:"region"`
	State        string           `json:"state"`
	City         string           `json:"city"`
	Address      string           `json:"address"`
	ZipCode      string           `json:"zipcode"`
	District     string           `json:"district"`
	Town         string           `json:"town"`
	AddressFlag  []string         `json:"address_flag"`
	TimeSlotList []TimeSlotEntity `json:"time_slot_list"`
	AddressType  []string         `json:"address_type"`
}

//String
func (a AddressEntity) String() string {
	return lib.ObjectToString(a)
}
