package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//DimensionEntity
type DimensionEntity struct{
	PackageLength int `json:"package_length"`
	PackageWidth int `json:"package_width"`
	PackageHeight int `json:"package_height"`
}

//String
func(p DimensionEntity)String()string{
	return lib.ObjectToString(p)
}