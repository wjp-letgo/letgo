package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//OptionEntity
type OptionEntity struct{
	Option string `json:"option"`
	Image TierImageEntity `json:"image"`
}

//String
func(s OptionEntity)String()string{
	return lib.ObjectToString(s)
}