package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//SizeEntity
type SizeEntity struct{
	SizeID string `json:"size_id"`
	Name string `json:"name"`
	DefaultPrice float32 `json:"default_price"`
}

//String
func(s SizeEntity)String()string{
	return lib.ObjectToString(s)
}