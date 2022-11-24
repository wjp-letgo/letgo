package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//DaysToShipLimitEntity
type DaysToShipLimitEntity struct{
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func(c DaysToShipLimitEntity)String()string{
	return lib.ObjectToString(c)
}