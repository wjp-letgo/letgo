package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//BrandEntity
type BrandEntity struct{
	BrandID int64 `json:"brand_id"`
	OriginalBrandName string `json:"original_brand_name"`
	DisplayBrandName string `json:"display_brand_name"`
}

//String
func(b BrandEntity)String()string{
	return lib.ObjectToString(b)
}