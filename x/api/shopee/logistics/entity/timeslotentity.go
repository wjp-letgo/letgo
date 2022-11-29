package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//TimeSlotEntity
type TimeSlotEntity struct {
	Date         int    `json:"date"`
	TimeText     string `json:"time_text"`
	PickupTimeID string `json:"pickup_time_id"`
}

//String
func (t TimeSlotEntity) String() string {
	return lib.ObjectToString(t)
}
