package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//BrandEntity
type BrandEntity struct{
	BrandID int64 `json:"brand_id"`
	OriginalBrandName string `json:"original_brand_name"`
}

//String
func(c BrandEntity)String()string{
	return lib.ObjectToString(c)
}