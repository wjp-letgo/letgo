package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//ItemMaxDimensionEntity
type ItemMaxDimensionEntity struct{
	Height float32 `json:"height"`
	Width float32 `json:"width"`
	Length float32 `json:"length"`
	Unit string `json:"unit"`
}

//String
func(i ItemMaxDimensionEntity)String()string{
	return lib.ObjectToString(i)
}