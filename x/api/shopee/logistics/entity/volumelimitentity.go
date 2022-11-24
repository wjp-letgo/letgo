package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//VolumeLimitEntity
type VolumeLimitEntity struct{
	ItemMaxVolume float32 `json:"item_max_volume"`
	ItemMinVolume float32 `json:"item_min_volume"`
}

//String
func(v VolumeLimitEntity)String()string{
	return lib.ObjectToString(v)
}