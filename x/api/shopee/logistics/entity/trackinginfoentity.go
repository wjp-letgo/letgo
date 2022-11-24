package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//TrackingInfoEntity
type TrackingInfoEntity struct{
	UpdateTime int `json:"update_time"`
	Description string `json:"description"`
}

//String
func(t TrackingInfoEntity)String()string{
	return lib.ObjectToString(t)
}